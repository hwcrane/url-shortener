package routes

import (
	"fmt"
	"net/http"
)

// HandleForm handles HTTP requests to the URL Shortener form.
func HandleForm(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST. If true, redirect to the "/shorten" route.
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/shorten", http.StatusSeeOther)
		return
	}

	// Serve the HTML form when the request method is not POST.
	w.Header().Set("Content-Type", "text/html")

	// Generate and write the HTML form to the response.
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>URL Shortener</title>
		</head>
		<body>
			<h2>URL Shortener</h2>
			<form method="post" action="/shorten">
				<input type="url" name="url" placeholder="Enter a URL" required>
				<input type="submit" value="Shorten">
			</form>
		</body>
		</html>
	`)
}
