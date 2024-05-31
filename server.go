package main

import (
	"fmt"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, r *http.Request) {
	fmt.Fprint(resp, m)
}

func main() {
	msgHandler := msg("Hello from Web Server in Go")
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", msgHandler)
}
