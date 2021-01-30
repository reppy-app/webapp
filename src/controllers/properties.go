package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"
)

// CreateProperty calls the API to create a property on database.
func CreateProperty(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	property, err := json.Marshal(struct {
		Name  string
		Price float64
	}{
		Name:  r.FormValue("name"),
		Price: price,
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/properties", config.APIURL)
	response, err := requests.AuthenticatedRequest(r, http.MethodPost, url, bytes.NewBuffer(property))
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
