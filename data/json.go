package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type APIError struct {
	Message string `json:"message"`
}

func FromJSON(r *http.Request, dest any) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(dest)
}

func RespondWithJSON(w http.ResponseWriter, payload any, code int) {
	w.Header().Add("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprint("Error marshalling JSON request: ", err)))
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, message string, code int) {
	if code > 499 {
		log.Print(message)
	}
	errorMsg := APIError{
		Message: message,
	}
	RespondWithJSON(w, errorMsg, code)
}
