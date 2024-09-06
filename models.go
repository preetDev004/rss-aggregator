// (OPTIONAL - You can avoid this file) - Personal preference
// This file makes SQLC types to snake-case types when responded in JSON.
// used to send only that much feilds which are required in response.

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
	for _, feed := range dbFeeds {
		feeds = append(feeds, dbFeedToFeed(feed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func dbFeedFollowToFeedFollow(dbFeedFollow db.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		FeedID:    dbFeedFollow.FeedID,
		UserID:    dbFeedFollow.UserID,
	}
}
func dbFeedFollowsToFeedFollows(dbFeedFollows []db.FeedFollow) []FeedFollow {
	var feedFollows []FeedFollow
	for _, feedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, dbFeedFollowToFeedFollow(feedFollow))
	}
	return feedFollows
}

type Post struct {
	ID          uuid.UUID      `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Title       string         `json:"title"`
	Description *string `json:"description"`
	PublishedAt time.Time      `json:"publushed_at"`
	Url         string         `json:"url"`
	FeedID      uuid.UUID      `json:"feed_id"`
}

func dbPostToPost(dbPost db.GetPostsForUserRow) Post {
	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: &dbPost.Description.String,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}
func dbPostsToPosts(dbPosts []db.GetPostsForUserRow) []Post {
	var posts []Post
	for _, post := range dbPosts {
		posts = append(posts, dbPostToPost(post))
	}
	return posts
}