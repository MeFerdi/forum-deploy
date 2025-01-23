package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3")

type SignUpErrors struct {
	NameError     string
	EmailError    string
	UsernameError string
	PasswordError string
	GeneralError  string
}

type SignUpData struct {
	Errors   SignUpErrors
	Email    string
	UserName string
}

var GlobalDB *sql.DB

func InitDB(database *sql.DB) {
	GlobalDB = database
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/signup.html"))
		tmpl.Execute(w, &SignUpData{})
		return
	}

	if r.Method == "POST" {
		data := SignUpData{
			UserName: r.FormValue("username"),
			Email:    r.FormValue("email"),
		}
		errors := SignUpErrors{}
		hasError := false

		// Validate name
		if data.UserName == "" {
			errors.UsernameError = "UserName is required"
			hasError = true
		}
	}
}
