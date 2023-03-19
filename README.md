## Create server ns and run server into that ns

```bash
sudo ip netns add serverns
# show all the network namespaces
ls -l /var/run/netns
# or
ip netns ls

# run server into the server namespace
sudo ip netns exec serverns ./server

./client
Requesting http://localhost:8080/serve
2023/03/18 12:47:56 Error Get "http://localhost:8080/serve": dial tcp 127.0.0.1:8080: connect: connection refused, getting localhost


# create client network ns and run client there
sudo ip netns add clientns

sudo ip netns exec client ./client
```

## Draw diagram to explain how we are going to create virt network interface to make the connection work

```bash
sudo ip netns exec clientns ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00

sudo ip netns exec serverns ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
```

## Create veth virtual ethernet network interface to conenct them

```bash
sudo ip link add pair-c type veth peer name pair-s

# run ip link to list all the network interfaces on the machine
ip link
```

## Move one end of the pair to server and another end to client ns

```bash
sudo ip link set pair-c netns clientns

# list network interfaces again from server and client to see new interface has been added
sudo ip netns exec serverns ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
15: pair-s@if16: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 46:11:c5:37:34:42 brd ff:ff:ff:ff:ff:ff link-netns clientns
```

## Assign IP address to the virtual interfaces

```bash
sudo ip netns exec clientns ip addr add 10.0.0.1/24 dev pair-c

sudo ip netns exec serverns ip addr add 10.0.0.2/24 dev pair-s
```

## Change the status to up

```
sudo ip netns exec serverns ip link set dev pair-s up
```

## Show the connectivity

```bash
# start server
sudo ip netns exec serverns ./server

# run client
sudo ip netns exec clientns ./client -server 10.0.0.2
Requesting http://10.0.0.2:8080/serve
Hello World!!!
```