package main

import (
	"fmt"
	"mmdist/pkg"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, r *http.Request) {
	fmt.Fprint(resp, m)
}

var types = []string{pkg.PersonalComputerType, pkg.NotebookType, pkg.ServerType, "monoblock"}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/html/main.html")
	})

	for _, typeName := range types {
		computer := pkg.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
