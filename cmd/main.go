package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.println(w, "Hello World")
}
