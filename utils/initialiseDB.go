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

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        post_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        likes INTEGER DEFAULT 0,
        dislikes INTEGER DEFAULT 0,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );
    CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts(user_id);
    CREATE INDEX IF NOT EXISTS idx_posts_post_at ON posts(post_at);
    `)
	if err != nil {
		fmt.Errorf("Failed to create posts table: %v", err)
		return nil
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        post_id INTEGER,
        user_id INTEGER,
        content TEXT NOT NULL,
        comment_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        likes INTEGER DEFAULT 0,
        dislikes INTEGER DEFAULT 0,
        FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );
    CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);
    CREATE INDEX IF NOT EXISTS idx_comments_user_id ON comments(user_id);
	`)
	if err != nil {
		fmt.Errorf("Failed to create comments table: %v", err)
		return nil
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT UNIQUE NOT NULL
    );
	`)
	if err != nil {
		fmt.Errorf("Failed to create categories table: %v", err)
		return nil
	}

	return db
}
