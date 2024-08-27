package main

import (
	"net/http"

	"github.com/lululululu5/blog-aggregator/auth"
	"github.com/lululululu5/blog-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Authorization Header wrong")
			return
		}

		user, err := cfg.DB.GetUserByAPI(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Could not get user")
			return
		}
		handler(w, r, user)
	}
}