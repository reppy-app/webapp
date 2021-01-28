package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CreateUser calls the API to create an user on database.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post("http://localhost:5000/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Body)
}
