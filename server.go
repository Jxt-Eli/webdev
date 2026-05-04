package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("server running on port:8080")

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n and this:\n %v", r.URL.Path, r)
}
