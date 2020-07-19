package main

import (
	"fmt"
	"net/http"
)

func handleDogRoute(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "dog route")
}

func handleCatRoute(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "cat route")
}

func main() {
	http.HandleFunc("/dog", handleDogRoute)
	http.HandleFunc("/cat", handleCatRoute)
	// Using the default ServeMux
	http.ListenAndServe(":8080", nil)
}
