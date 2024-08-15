package main

import (
	"fmt"
	"net/http"

	"github.com/preetDev004/rss-aggregator/auth"
	"github.com/preetDev004/rss-aggregator/db"
)

type authHandler func(http.ResponseWriter, *http.Request, db.User)

// closer for middleware
func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("couldn't find the user: %v", err))
			return
		}
		handler(w, r, user)
	}
}