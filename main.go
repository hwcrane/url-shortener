package main

import (
	"fmt"
	"net/http"
	"url-shortener/routes"
)

func main() {
	http.HandleFunc("/", routes.HandleForm)
	http.HandleFunc("/shorten", routes.HandelShorten)
	http.HandleFunc("/short/", routes.HandleRedirect)

	fmt.Println("URL Shortener is running on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
