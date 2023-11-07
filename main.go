package main

import (
	"fmt"
	"net/http"
	"url-shortener/routes"
)

func main() {
	// Define HTTP routes and their associated handlers.

	// HandleForm serves the URL Shortener form at the root path.
	http.HandleFunc("/", routes.HandleForm)

	// HandelShorten handles POST requests to shorten a URL.
	http.HandleFunc("/shorten", routes.HandelShorten)

	// HandleRedirect redirects short URLs to their original counterparts.
	http.HandleFunc("/short/", routes.HandleRedirect)

	// Print a message indicating that the URL Shortener is running.
	fmt.Println("URL Shortener is running on localhost:8080")

	// Start the HTTP server to listen on localhost and port 8080.
	http.ListenAndServe("localhost:8080", nil)
}
