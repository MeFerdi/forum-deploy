package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var GlobalDB *sql.DB

func InitialiseDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		log.Fatal(err)
	}

	GlobalDB = db

	// Create Users table
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT UNIQUE,
            username TEXT UNIQUE,
            password TEXT
        );
		CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
    	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
    `)
	if err != nil {
		fmt.Errorf("Failed to create users table: %v", err)
		return nil
	}

	return db
}
