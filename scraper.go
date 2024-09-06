package main

import (
	"context"
	"strings"
	"os"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/preetDev004/rss-aggregator/db"
)

func startScraping(db *db.Queries, concurruncy int, timeBetweenRequests time.Duration) {
	log.Printf("Scraping on %v go-routines every %s", concurruncy, timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	// every 'timeBetweenRequests' there will be a value passed in the ticker.C channel
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurruncy))
		if err != nil {
			log.Println("error fetching feeds", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapFeed(db, wg, feed)
		}
		wg.Wait()
	}
}
func scrapFeed(database *db.Queries, wg *sync.WaitGroup, feed db.Feed) {
	defer wg.Done()

	_, err := database.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched", err)
	}
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("failed to fetch feed", err)
		return
	}
	data := createPostsSlice(rssFeed.Channel.Item, feed.ID)

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Prepare the copy command
	_, err = conn.CopyFrom(
		context.Background(),
		pgx.Identifier{"posts"}, // table name
		[]string{"id", "created_at", "updated_at", "title", "description", "url", "published_at", "feed_id"}, // column names
		pgx.CopyFromSlice(len(data), func(i int) ([]interface{}, error) {
			return []interface{}{data[i].ID, data[i].CreatedAt, data[i].UpdatedAt, data[i].Title, data[i].Description, data[i].Url, data[i].PublishedAt, data[i].FeedID}, nil
		}),
	)
	if err != nil && !strings.Contains(err.Error(), "duplicate key"){
		log.Printf("Couldn't create the posts %v", err)
	}

	log.Printf("Feed %s Collected, %v Posts Found", feed.Name, len(rssFeed.Channel.Item))

}