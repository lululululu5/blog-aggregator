package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lululululu5/blog-aggregator/internal/database"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}

	type response struct {
		Feed      Feed      `json:"feed"`
		FeedFollow FeedFollow `json:"feed_follow"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		FeedID: feed.ID,
		UserID: user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return 
	}

	resp := response{
		Feed:      databaseFeedToFeed(feed),
		FeedFollow: databaseFeedFollowToFeedFoolow(feedFollow),
	}

	respondWithJSON(w, http.StatusOK, resp)
}

func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds)) // Important to transform the database model to the correct API representation for Flexibility, Security, Data transformation, Versioning and Seperation of concerns. 
}

func (cfg *apiConfig) handlerGetNextFeedsToFetch(w http.ResponseWriter, r *http.Request) {
	var numNextFeeds int32 = 10
	feeds, err := cfg.DB.GetNextFeedsToFetch(r.Context(), numNextFeeds)
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Couldn't fetch feeds")
		return 
	}
	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}

// This could also be a method instead of a handler. But this handler will be with the feedID in the header
func (cfg *apiConfig) handlerMarkFeedFetched(w http.ResponseWriter, r *http.Request) {
	//Add feedID to url
	feedID, err := uuid.Parse(r.PathValue("feedID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not extract Feed ID")
		return 
	}
	
	err = cfg.DB.MarkFeedFetched(r.Context(), database.MarkFeedFetchedParams{
		ID: feedID,
		UpdatedAt: time.Now().UTC(),
		LastFetchedAt: sql.NullTime{
			Time: time.Now().UTC(),
			Valid: true,
		},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not mark Feed as fetched.")
		return
	}

	respondWithJSON(w, http.StatusNoContent, struct{}{})
}

