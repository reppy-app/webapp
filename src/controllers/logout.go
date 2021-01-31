package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Logout remove user data saved on browser.
func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", 302)
}
