package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorAPI represents the response of error from API.
type ErrorAPI struct {
	Error string `json:"error"`
}

// JSON returns a JSON response to the request.
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// HandleStatusCodeError handle the requests with status code 400 or  higher.
func HandleStatusCodeError(w http.ResponseWriter, r *http.Response) {
	var err ErrorAPI
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
