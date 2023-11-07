package database

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
)

// DB is a global database connection variable.
var DB *sql.DB

// init initializes the database connection.
func init() {
	// Open an SQLite database connection and store it in the global DB variable.
	DB, _ = sql.Open("sqlite3", "data.db")
	DB.Exec("CREATE TABLE IF NOT EXISTS urls( key      VARCHAR(128) NOT NULL, url     VARCHAR(255) NOT NULL, PRIMARY KEY (`key`) )")
}

// AddKey inserts a new URL and its associated key into the database.
func AddKey(url string, key string) {
	println(DB)

	// Prepare an SQL statement for insertion.
	statement, _ := DB.Prepare("INSERT INTO urls (key, url) VALUES (?, ?)")

	// Execute the prepared statement with the provided key and URL.
	statement.Exec(key, url)
}

// GetKey retrieves the URL associated with a given key from the database.
func GetKey(key string) string {
	// Define the SQL query statement.
	statement := "SELECT url FROM urls WHERE key = ?"

	// Execute the query and retrieve a single row.
	rows := DB.QueryRow(statement, key)

	var temp string

	// Scan the result into the temp variable.
	err := rows.Scan(&temp)

	// Check if no rows were found, and return an empty string in that case.
	if errors.Is(err, sql.ErrNoRows) {
		return ""
	}

	// Return the URL associated with the provided key.
	return temp
}

// GenerateShortKey generates a random short key for shortening URLs.
func GenerateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	// Create a byte slice to store the random key.
	shortKey := make([]byte, keyLength)

	// Generate a random key by selecting characters from the charset.
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}

	// Convert the byte slice to a string and return the generated short key.
	return string(shortKey)
}
