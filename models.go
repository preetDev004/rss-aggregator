// (OPTIONAL - You can avoid this file) - Personal preference
// This file makes SQLC types to snake-case types when responded in JSON.
package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/preetDev004/rss-aggregator/db"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func dbUserToUser(user db.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func dbFeedToFeed(feed db.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}
func dbFeedsToFeeds(dbFeeds []db.Feed) []Feed {
	var feeds []Feed
	for _, feed := range dbFeeds{
		feeds = append(feeds, dbFeedToFeed(feed))
	}
	return feeds
}

type FeedFollow struct{
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}
func dbFeedFollowToFeedFollow(dbFeedFollow db.FeedFollow) FeedFollow{
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		FeedID:    dbFeedFollow.FeedID,
		UserID:    dbFeedFollow.UserID,
	}
}
func dbFeedFollowsToFeedFollows(dbFeedFollows []db.FeedFollow) []FeedFollow{
	var feedFollows []FeedFollow
	for _, feedFollow := range dbFeedFollows{
		feedFollows = append(feedFollows, dbFeedFollowToFeedFollow(feedFollow))
	}
	return feedFollows
}