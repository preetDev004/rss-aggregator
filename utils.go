package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/preetDev004/rss-aggregator/db"
)

type apiConfig struct {
	DB *db.Queries
}

func connectToDB(dbURL string) *sql.DB {
	connection, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to database")
	return connection
}

func closeDB(db *sql.DB) {
	log.Println("Closing Connection to the Database.")
	db.Close()
}

func createPostsSlice(items []RSSItem, feedID uuid.UUID) ([]db.Post) {
	data := make([]db.Post, len(items))

	for i, item := range items {
		desc := sql.NullString{
			String: item.Description,
		}
		if item.Description == "" {
			desc.Valid = true
		}
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Couldn't parse the published date %v with err %v", item.PubDate, err)
			continue
		}

		data[i].ID = uuid.New()
		data[i].CreatedAt = time.Now().UTC()
		data[i].UpdatedAt = time.Now().UTC()
		data[i].Title = item.Title
		data[i].Description = desc
		data[i].Url = item.Link
		data[i].PublishedAt = pubDate
		data[i].FeedID = feedID
	}
	return data
}
