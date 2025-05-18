package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gerard-007/bootdotdev/internal/database"
	"github.com/google/uuid"
	"github.com/go-chi/chi"
)

func (apiConfig *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithJsonError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		FeedID: params.FeedID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithJsonError(w, 400, fmt.Sprintf("Couldnt create feed follow %v", err))
		return
	}
	responseWithJson(w, 200, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiConfig *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	follows, err := apiConfig.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithJsonError(w, 400, fmt.Sprintf("Couldnt get feed follows %v", err))
		return
	}
	responseWithJson(w, 200, databaseFeedFollowToFeedFollows(follows))
}

func (apiConfig *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithJsonError(w, 400, fmt.Sprintf("Error parsing feed follow ID: %v", err))
		return
	}
	err = apiConfig.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithJsonError(w, 400, fmt.Sprintf("Couldnt delete feed follow %v", err))
		return
	}
	responseWithJson(w, 200, map[string]string{"message": "Feed follow deleted"})
}