package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello world")
}

func main() {
	http.HandleFunc("/", handler)
}
