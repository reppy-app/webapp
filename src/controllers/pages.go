package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// LoadLoginPage load login page.
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecuteTemplate(w, "signin.html", nil)
}

// LoadSignupPage load signin page.
func LoadSignupPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signup.html", nil)
}

// LoadHome load home page.
func LoadHome(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/properties", config.APIURL)
	response, err := requests.AuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}

	var properties []models.Property
	if err = json.NewDecoder(response.Body).Decode(&properties); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	utils.ExecuteTemplate(w, "home.html", struct {
		Properties []models.Property
		UserID     uint64
	}{
		Properties: properties,
		UserID:     userID,
	})
}

// LoadEditPropertyPage load page to edit property.
func LoadEditPropertyPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	propertyID, err := strconv.ParseUint(params["propertyId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/properties/%d", config.APIURL, propertyID)
	response, err := requests.AuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}

	var property models.Property
	if err := json.NewDecoder(response.Body).Decode(&property); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "edit-property.html", property)
}
