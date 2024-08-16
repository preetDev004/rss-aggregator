package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/preetDev004/rss-aggregator/db"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON: %v", err))
		return
	}
	if len(params.Url) == 0 {
		respondWithError(w, 400, "Please enter a valid URL")
		return
	}
	if len(params.Name) <= 2 || len(params.Name) > 15 {
		respondWithError(w, 400, fmt.Sprintf("Please enter a name with more than %v characters and less than %v characters!", MIN_CHARS, MAX_CHARS))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error Creating the feed: %v", err))
		return
	}
	respondWithJSON(w, 201, dbFeedToFeed(feed))
}

func (apiCfg *apiConfig) handleGetAllFeeds(w http.ResponseWriter, r *http.Request){
	
	dbFeeds, err:=apiCfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("No feeds found: %v", err))
		return
	}
	var feeds []Feed
	for _, f := range dbFeeds{
		feeds = append(feeds, dbFeedToFeed(f))
	}
	respondWithJSON(w, 200, feeds)

}