package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func responseWithJsonError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Printf("Responding with 5xx error: %s", message)
	}
	type ErrorResponse struct {
		Error string `json:"error"`
	}
	responseWithJson(w, code, ErrorResponse{
		Error: message,
	})
}