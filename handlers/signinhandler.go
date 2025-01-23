package handlers

import "net/http"

type SignInData struct {
	Username      string
	UsernameError string
	PasswordError string
	GeneralError  string
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {

}