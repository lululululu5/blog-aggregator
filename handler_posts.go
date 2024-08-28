package main

import (
	"context"
	"net/http"

	"github.com/lululululu5/blog-aggregator/internal/database"
)

func (cfg *apiConfig) handlerGetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.DB.GetPostsByUser(context.Background(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't gather posts by user")
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}