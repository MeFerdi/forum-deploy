package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
)

type SignUpData struct {
	Errors   SignUpErrors
	Name     string
	Email    string
	Username string
}

var GlobalDB *sql.DB

func InitDB(database *sql.DB) {
	GlobalDB = database
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/signup.html"))
		tmpl.Execute(w, &SignUpData)
		return
	}
}
