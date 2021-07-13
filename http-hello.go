package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello! %s <h1>", r.URL)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8800", nil)
}
