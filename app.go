package main

import (
	"io"
	"net/http"
)

func main() {
	// Create a basic http example to demonstrate example
	http.Handle("/", http.HandlerFunc(ExampleHandler))
	http.ListenAndServe(":8080", nil)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	// This handler will run for all types of HTTP request, but we can use r.Method to
	// determine which method is being used and validate the request based on this.
	if r.Method == http.MethodGet {
		io.WriteString(w, "Great!!! You get it")
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "Great!!! You post it")
	} else {
		io.WriteString(w, "This is a "+r.Method+" request")
	}
}
