package web

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var (
	ErrManNotFound = errors.New("man not found")
	ErrJsonEncode  = errors.New("json encode failed")
	ErrBadRequest  = errors.New("Err Bad Request")
)

// responseError it`s structure sending error
type responseError struct {
	StatusCode  int    `json:"statusCode"`
	Description string `json:"description"`
}

// SendError this universal sender error
func (dh *DataHandler) SendError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	if sendErr := json.NewEncoder(w).Encode(responseError{
		StatusCode:  statusCode,
		Description: err.Error(),
	}); sendErr != nil {
		log.Printf("response send error : %s", sendErr)
	}
}
