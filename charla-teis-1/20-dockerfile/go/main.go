package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	println("GET /hello")
	fmt.Fprintf(w, "<h1>Hello world!</h1>\n")
}

func main() {
	println("Started app...")
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
