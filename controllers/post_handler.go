package controllers

import (
	"html/template"
	"net/http"
)

func PostHandler(writer http.ResponseWriter, reader *http.Request) {
	temp1, err := template.ParseFiles("templates/post.html")
	if err != nil {
		http.Error(writer, "Error loading template", 500)
		return
	}
	err = temp1.Execute(writer, "")
	if err != nil {
		http.Error(writer, "Error loading template", 500)
	}
}
