package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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
	feedFollow, err := apiCfg.DB.CreateFeedFollows(r.Context(), db.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating the feed follow: %v", err))
		return
	}
	respondWithJSON(w, 201, dbFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handleGetUserFeedFollows(w http.ResponseWriter, r *http.Request, user db.User){
	FeedFollows, err := apiCfg.DB.GetUserFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't find any feed follows: %v", err))
		return
	}
	if FeedFollows == nil{
		respondWithJSON(w, 200, successful{Msg: "You Don't follow any feeds right now!"})
		return
	}

	respondWithJSON(w, 200, dbFeedFollowsToFeedFollows(FeedFollows))

}
func (apiCfg *apiConfig) handleDeleteUserFeedFollow(w http.ResponseWriter, r *http.Request, user db.User){
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowId, err:=uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}
	_, err = apiCfg.DB.DeleteUserFeedFollow(r.Context(),db.DeleteUserFeedFollowParams{
		FeedID: feedFollowId,
		UserID: user.ID,
	})
	
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete the feed follows: %v", err))
		return
	}
	
	respondWithJSON(w, 200, successful{Msg:"Unfollowed the feed, successfully!"})
}