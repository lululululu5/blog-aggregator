package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lululululu5/blog-aggregator/internal/database"
)

func (cfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return 
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		FeedID: params.FeedID,
		UserID: user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return 
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFoolow(feedFollow))
}

func (cfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID, err := uuid.Parse(r.PathValue("feedFollowID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode Feed Follow ID")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID : feedFollowID,
		UserID: user.ID,
	} )
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not delete Feed Follow")
		return
	}

	respondWithJSON(w, http.StatusNoContent, "")
}

func (cfg *apiConfig) handlerGetFeedFollowUser(w http.ResponseWriter, r *http.Request, user database.User) {
	FeedFollows, err := cfg.DB.GetFeedFollowsUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get User's followed feeds")
		return 
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(FeedFollows))
}