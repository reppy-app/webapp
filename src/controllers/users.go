package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/responses"
)

// CreateUser calls the API to create an user on database.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:5000/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}
	responses.JSON(w, response.StatusCode, nil)
}
