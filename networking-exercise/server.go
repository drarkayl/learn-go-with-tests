package main

import (
	"fmt"
	"net/http"
)

// The handler function for our HTTP server.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// The main function that starts the HTTP server.
func main() {
	// Register the handler for the "/" route.
	http.HandleFunc("/", helloHandler)

	// Start the server on port 8081.
	fmt.Println("Server is starting on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
