package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// before init the server listener same as express js
	http.HandleFunc("/", handleGetRequest)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Stating server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	fmt.Println(r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Helldsado!")
}
