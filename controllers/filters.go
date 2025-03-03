package controllers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	"forum/utils"
)

func CreatedPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Check session
	userID, err := validateUserSession(w, r)
	if err != nil {
		return
	}

	// Fetch posts
	posts, err := fetchUserPostsForPosts(userID)
	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	// Fetch all users
	users, err := getAllUsers()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
	}

	if err := renderCreatedTemplateForPosts(w, posts, users, userID); err != nil {
		log.Printf("Error rendering template: %v", err)
		return
	}
}

func LikedPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Check session
	userID, err := validateUserSession(w, r)
	if err != nil {
		return
	}

	// Fetch posts
	posts, err := fetchUserPostsForLikes(userID)
	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	// Fetch all users
	users, err := getAllUsers()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
	}

	if err := renderCreatedTemplateForLikes(w, posts, users, userID); err != nil {
		log.Printf("Error rendering template: %v", err)
		return
	}
}

// getAllUsers fetches all users from the database
func getAllUsers() ([]utils.User, error) {
	rows, err := utils.GlobalDB.Query(`
		SELECT id, username, profile_pic 
		FROM users
		ORDER BY username
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []utils.User
	for rows.Next() {
		var user utils.User
		err := rows.Scan(&user.ID, &user.UserName, &user.ProfilePic)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func validateUserSession(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return "", err
	}

	userID, err := utils.ValidateSession(utils.GlobalDB, cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return "", err
	}

	return userID, nil
}

func fetchUserPostsForPosts(userID string) ([]utils.Post, error) {
	rows, err := utils.GlobalDB.Query(`
        SELECT p.id, p.user_id, p.title, p.content, p.imagepath, p.post_at, p.likes, p.dislikes, p.comments,
               u.username, u.profile_pic, c.id AS category_id, c.name AS category_name
        FROM posts p
        JOIN users u ON p.user_id = u.id
        LEFT JOIN post_categories pc ON p.id = pc.post_id
        LEFT JOIN categories c ON pc.category_id = c.id
        WHERE p.user_id = ?
        ORDER BY p.post_at DESC
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	postMap := make(map[int]utils.Post)
	var postTime time.Time
	for rows.Next() {
		var post utils.Post
		var categoryID sql.NullInt64
		var categoryName sql.NullString
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.ImagePath,
			&postTime,
			&post.Likes,
			&post.Dislikes,
			&post.Comments,
			&post.Username,
			&post.ProfilePic,
			&categoryID,
			&categoryName,
		)
		if err != nil {
			return nil, err
		}
		post.PostTime = FormatTimeAgo(postTime.Local())

		postMap[post.ID] = post
	}

	var posts []utils.Post
	for _, post := range postMap {
		posts = append(posts, post)
	}

	return posts, nil
}

func fetchUserPostsForLikes(userID string) ([]utils.Post, error) {
	rows, err := utils.GlobalDB.Query(`
        SELECT p.id, p.user_id, p.title, p.content, p.imagepath, p.post_at, p.likes, p.dislikes, p.comments,
               u.username, u.profile_pic, c.id AS category_id, c.name AS category_name
        FROM posts p
        JOIN users u ON p.user_id = u.id
        LEFT JOIN post_categories pc ON p.id = pc.post_id
        LEFT JOIN categories c ON pc.category_id = c.id
        JOIN reaction r ON p.id = r.post_id
        WHERE r.user_id = ? AND r.like = 1 OR r.like = 0
        ORDER BY p.post_at DESC
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	postMap := make(map[int]utils.Post)
	var postTime time.Time
	for rows.Next() {
		var post utils.Post
		var categoryID sql.NullInt64
		var categoryName sql.NullString
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.ImagePath,
			&postTime,
			&post.Likes,
			&post.Dislikes,
			&post.Comments,
			&post.Username,
			&post.ProfilePic,
			&categoryID,
			&categoryName,
		)
		if err != nil {
			return nil, err
		}
		post.PostTime = FormatTimeAgo(postTime.Local())

		postMap[post.ID] = post
	}

	var posts []utils.Post
	for _, post := range postMap {
		posts = append(posts, post)
	}

	return posts, nil
}

func renderCreatedTemplateForPosts(w http.ResponseWriter, posts []utils.Post, users []utils.User, userID string) error {
	tmpl, err := template.ParseFiles("templates/created.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return err
	}

	data := struct {
		Posts  []utils.Post
		Users  []utils.User
		UserID string
	}{
		Posts:  posts,
		Users:  users,
		UserID: userID,
	}

	return tmpl.Execute(w, data)
}

func renderCreatedTemplateForLikes(w http.ResponseWriter, posts []utils.Post, users []utils.User, userID string) error {
	tmpl, err := template.ParseFiles("templates/liked.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return err
	}

	data := struct {
		Posts  []utils.Post
		Users  []utils.User
		UserID string
	}{
		Posts:  posts,
		Users:  users,
		UserID: userID,
	}

	return tmpl.Execute(w, data)
}
