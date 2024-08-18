package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/preetDev004/rss-aggregator/db"
)

func startScraping(db *db.Queries, concurruncy int, timeBetweenRequests time.Duration){
	log.Printf("Scraping on %v go-routines every %s", concurruncy, timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	// every 'timeBetweenRequests' there will be a value passed in the ticker.C channel
	for ; ; <-ticker.C{
		feeds, err:=db.GetNextFeedsToFetch(context.Background(), int32(concurruncy))
		if err!= nil{
			log.Println("error fetching feeds",err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds{
			wg.Add(1)
			go scrapFeed(db, wg, feed)
		}
		wg.Wait()
	}
}
func scrapFeed(db *db.Queries,wg *sync.WaitGroup, feed db.Feed ){
	defer wg.Done()

	_ , err:=db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched",err)
	}
	rssFeed, err:=urlToFeed(feed.Url)
	if err != nil {
		log.Println("failed to fetch feed",err)
		return
	}
	for _, item := range rssFeed.Channel.Item{
		log.Println("Found Post:",item.Title, "On feed", feed.Name)
	}
	log.Printf("Feed %s Collected, %v Posts Found", feed.Name, len(rssFeed.Channel.Item))

	// update the database with the updated feeds
	
}