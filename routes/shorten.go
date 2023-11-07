package routes

import (
	"fmt"
	"net/http"
	"url-shortener/database"
)

// HandelShorten is the handler for shortening a URL.
func HandelShorten(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests for shortening a URL.
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the original URL from the request form.
	originalUrl := r.FormValue("url")

	// Check if the URL parameter is missing in the form.
	if originalUrl == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// Generate a new short key and associate it with the original URL in the database.
	shortKey := database.GenerateShortKey()
	database.AddKey(originalUrl, shortKey)

	// Construct the shortened URL with the generated short key.
	shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	// Return an HTML page with information about the original and shortened URLs.
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

	// Write the HTML response to the client.
	fmt.Fprintf(w, responseHTML)
}
