package utils

import (
	"database/sql"
	"time"
)

type User struct {
	ID         string
	UserName   string
	Email      string
	Password   string
	ProfilePic sql.NullString
}

type Post struct {
	ID           int
	UserID       string
	Title        string
	Content      string
	ImagePath    string
	PostTime     string
	Likes        int
	Dislikes     int
	Comments     int
	Username     string
	ProfilePic   sql.NullString
	UserReaction *int
	CategoryID   *int
	CategoryName *string
}

type Comment struct {
	ID          int
	PostID      int
	UserID      string
	Username    string
	Content     string
	CommentTime time.Time
	Likes       int
	Dislikes    int
	ProfilePic  sql.NullString
}

type Category struct {
	ID   int
	Name string
}

type Session struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
}

type ErrorPageData struct {
	Code    int
	Message string
}

type PageData struct {
	IsLoggedIn    bool
	Posts         []Post
	CurrentUserID string
	Users []User
}

type Notification struct {
    ID                 int
    Type              string    // "like", "dislike", "comment"
    PostID            int       // ID of the affected post
    ActorName         string    // Username of person who performed action
    ActorProfilePic   sql.NullString // Profile picture of actor
    CreatedAt         time.Time // When notification was created
    CreatedAtFormatted string   // Formatted time string (e.g., "2 hours ago")
}
