package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// LoadLoginPage load login page.
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signin.html", nil)
}

// LoadSignupPage load signin page.
func LoadSignupPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signup.html", nil)
}
