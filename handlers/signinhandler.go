package handlers

import (
	"html/template"
	"net/http"
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
	}
}
