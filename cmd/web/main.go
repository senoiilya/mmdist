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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/html/main.html")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
