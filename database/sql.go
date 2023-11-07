package database

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
)

var DB *sql.DB

func init() {
	DB, _ = sql.Open("sqlite3", "data.db")

}

func AddKey(url string, key string) {
	println(DB)
	statement, _ := DB.Prepare("INSERT INTO urls (key, url) VALUES (?, ?)")
	statement.Exec(key, url)
}

func GetKey(key string) string {
	statement := "SELECT url FROM urls WHERE key = ?"
	rows := DB.QueryRow(statement, key)
	var temp string
	err := rows.Scan(&temp)
	if errors.Is(err, sql.ErrNoRows) {
		return ""
	}
	return temp
}

func GenerateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
