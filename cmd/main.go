package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Print("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
