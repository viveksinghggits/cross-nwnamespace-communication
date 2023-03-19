package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/serve", handleServe)
	http.ListenAndServe(":8080", r)
}

func handleServe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling serve")
	w.Write([]byte(fmt.Sprintf("%s\n", "Hello World!!!")))
}
