package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gerard-007/bootdotdev/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json: "username"`
		Email    string `json: "email"`
		Password string `json: "password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithJsonError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Email:     params.Email,
		Password:  params.Password,
		Username:  params.Username,
		CreatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
	})
	if err != nil {
		responseWithJsonError(w, 400, fmt.Sprintf("Couldnt create user %v", err))
		return
	}
	responseWithJson(w, 200, databaseUserToUser(user))
}


func (apiConfig *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) { 
	responseWithJson(w, 200, databaseUserToUser(user))
}


func (apiConfig *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) { 
	posts, err := apiConfig.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		responseWithJsonError(w, 400, fmt.Sprintf("Couldnt get posts for user %v", err))
		return
	}
	responseWithJson(w, 200, databasePostsToPosts(posts))
}

