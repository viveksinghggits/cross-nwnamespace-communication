package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var serverAdd string
	flag.StringVar(&serverAdd, "server", "localhost", "Address of the server")
	flag.Parse()

	fmt.Printf("Requesting %s\n", fmt.Sprintf("http://%s:8080/serve", serverAdd))
	uri := fmt.Sprintf("http://%s:8080/serve", serverAdd)
	resp, err := http.Get(uri)
	if err != nil {
		log.Printf("Error %s, getting %s", err.Error(), serverAdd)
		return
	}

	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error %s reading response data", err.Error())
		return
	}

	fmt.Println(string(respBody))
}
