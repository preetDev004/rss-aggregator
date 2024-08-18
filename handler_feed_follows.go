package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/preetDev004/rss-aggregator/db"
)

func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON: %v", err))
		return
	}
	isExists, err := apiCfg.DB.IsFeedExists(r.Context(), params.FeedID)
	if err != nil || !isExists{
		respondWithError(w, 400, fmt.Sprintf("Couldn't find the feed with id: %v",params.FeedID))
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollows(r.Context(), db.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error Creating the feed follow: %v", err))
		return
	}
	respondWithJSON(w, 201, dbFeedFollowToFeedFollow(feedFollow))
}
