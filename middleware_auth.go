package main

import (
	"fmt"
	"net/http"

	"github.com/Gerard-007/bootdotdev/internal/auth"
	"github.com/Gerard-007/bootdotdev/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(next authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithJsonError(w, 403, fmt.Sprintf("unauthorized: %s", err))
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithJsonError(w, 403, fmt.Sprintf("unauthorized: %s", err))
			return
		}
		next(w, r, user)
	}
}
