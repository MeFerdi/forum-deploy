package handlers

import (
	"database/sql"
	"html/template"
	"log"
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
	_, err := r.Cookie("session_token")
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("templates/signin.html")
	if err != nil {
		utils.RenderErrorPage(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		log.Printf("Error loading template: %v", err)
		return
	}

	if r.Method == "GET" {
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		data := SignInData{}

		username := r.FormValue("username")
		password := r.FormValue("password")

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
		if err != nil {
			data := struct {
				GeneralError string
				Username     string
			}{
				GeneralError: "Invalid username or password",
				Username:     username,
			}
			tmpl, _ := template.ParseFiles("templates/signin.html")
			tmpl.Execute(w, data)
			if err != sql.ErrNoRows {
				log.Printf("Error querying database: %v", err)
			}
			return
		}

		if !utils.CheckPasswordsHash(password, user.Password) {
			data.GeneralError = "Invalid username or password"
			data.Username = username
			tmpl.Execute(w, data)
			return
		}

		sessionToken, err := utils.CreateSession(GlobalDB, user.ID)
		if err != nil {
			utils.RenderErrorPage(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   24 * 60 * 60,
		})

		log.Printf("Set session cookie: %s", sessionToken)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
