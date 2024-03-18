package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// StatusResponse represents a general status response
type StatusResponse struct {
	Status string `json:"status"`
}

// NewErrorResponse sends an error response with a message and logs the error
func NewErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	logrus.Error(message)
	respondJSON(w, statusCode, ErrorResponse{Message: message})
}

// NewStatusResponse sends a general status response
func NewStatusResponse(w http.ResponseWriter, statusCode int, status string) {
	respondJSON(w, statusCode, StatusResponse{Status: status})
}

// respondJSON marshals the payload to JSON and writes it to the response writer
func respondJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func decodeJSONFromBody(r *http.Request, dst interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, dst)
	if err != nil {
		return err
	}

	return nil
}
