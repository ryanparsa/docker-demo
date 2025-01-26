package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! This is a Go app.")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on 0.0.0.0:8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
