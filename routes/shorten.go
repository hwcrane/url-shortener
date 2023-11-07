package routes

import (
	"fmt"
	"net/http"
	"url-shortener/database"
)

// HandelShorten is the handler for shortening a URL
func HandelShorten(w http.ResponseWriter, r *http.Request) {
	// Only works for post requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the url
	originalUrl := r.FormValue("url")
	if originalUrl == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// Generate the new key and assign it
	shortKey := database.GenerateShortKey()
	database.AddKey(originalUrl, shortKey)

	// The new shortened URL
	shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	// Return a html page with the new shortened URL in it
	w.Header().Set("Content-Type", "text/html")
	responseHTML := fmt.Sprintf(`
	<h2>URL Shortener</h2>
	<p>Original URL: %s</p>
	<p>Shortened URL: <a href="%s">%s</a></p>
	<form method="post" action="/shorten">
	<input type="text" name="url" placeholder="Enter a URL">
	<input type="submit" value="Shorten">
	</form>`, originalUrl, shortenedURL, shortenedURL,
	)
	fmt.Fprintf(w, responseHTML)
}
