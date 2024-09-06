package main

import (
	"fmt"
	"net/http"
	"github.com/preetDev004/rss-aggregator/db"
)

func (apiCfg *apiConfig) handleGetPosts(w http.ResponseWriter, r *http.Request, user db.User){
	userPosts, err := apiCfg.DB.GetPostsForUser(r.Context(), db.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't find the posts for user %v", err))
		return
	}
	if userPosts == nil {
		respondWithJSON(w, 200, successful{Msg: "There are no posts to show you."})
	}
	respondWithJSON(w, 200, dbPostsToPosts(userPosts))
} 