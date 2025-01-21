package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/controllers"
	"forum/utils"
)

func main() {
	db := utils.InitialiseDB()
	if db != nil {
		defer db.Close()
	}
	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/post", controllers.PostHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server opened at port 3000...http://localhost:3000/")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println("Failed to ope server")
		return
	}
}
