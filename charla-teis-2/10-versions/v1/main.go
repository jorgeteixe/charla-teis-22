package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	println("(v1) GET /")
	fmt.Fprintf(w, "<h1>Hello world from v1!</h1>\n")
}

func main() {
	println("Started app...")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}
