package main

import (
	"fmt"
	"net/http"
)

func main() {
	
	// Registering a Request Handler
	
	http.HandleFunc("/Dhik", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello you've requested: %s\n", r.URL.Path)
	})
	
	// Setting up an HTTP listner on port 8080
	
	http.ListenAndServe(":8080", nil)
}
