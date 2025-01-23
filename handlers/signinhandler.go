package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"forum/utils"
)

type SignInData struct {
	Username      string
	UsernameError string
	PasswordError string
	GeneralError  string
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/signin.html"))

	if r.Method == "GET" {
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		data := SignInData{}

		// get form values
		username := r.FormValue("username")
		password := r.FormValue("password")

		// validate input
		if username == "" {
			data.UsernameError = "Username is required"
		}

		if password == "" {
			data.PasswordError = "password is required"
		}

		if data.UsernameError != "" || data.PasswordError != "" {
			data.Username = username
			tmpl.Execute(w, data)
			return
		}

		var user utils.User
		err := GlobalDB.QueryRow(`
			SELECT id, password
			FROM users
			WHERE username = ?
		`, username).Scan(&user.ID, &user.Password)
		// Handle database errors
		if err != nil {
			if err == sql.ErrNoRows {
				data.GeneralError = "Invalid username or password"
			} else {
				data.GeneralError = "An error occurred. Please try again later."
			}
			data.Username = username // Preserve the username input
			tmpl.Execute(w, data)
			return
		}

		// Verify password
		if !utils.CheckPasswordsHash(password, user.Password) {
			data.GeneralError = "Invalid username or password"
			data.Username = username // Preserve the username input
			tmpl.Execute(w, data)
			return
		}

	}
}
