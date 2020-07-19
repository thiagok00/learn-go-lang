package main

import (
	"fmt"
	"net/http"
)

type myHTTPHandler int

func (m myHTTPHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	switch req.URL.Path {
	case "/about":
		fmt.Fprint(w, "about")
	case "/":
		fmt.Fprint(w, "hello http")
	default:
		w.WriteHeader(404)
	}

}

func main() {
	var d myHTTPHandler
	http.ListenAndServe(":8080", d)
}
