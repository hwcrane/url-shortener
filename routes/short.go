package routes

import (
	"net/http"
	"url-shortener/database"
)

// HandleRedirect handles HTTP requests to redirect to the original URL using a shortened key.
func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	// Extract the shortened key from the request URL.
	shortKey := r.URL.Path[len("/short/"):]

	// Check if the shortened key is missing.
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	// Retrieve the original URL associated with the shortened key from the database.
	originalURL := database.GetKey(shortKey)

	// Check if the shortened key was not found in the database.
	if originalURL == "" {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	// Redirect the client to the original URL with a 301 (Moved Permanently) status code.
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
