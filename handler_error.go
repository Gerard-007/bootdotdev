package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"
	responseWithJsonError(w, code, message)
}
