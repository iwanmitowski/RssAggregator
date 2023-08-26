package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/iwanmitowski/RssAggregator/internal/database"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf(
		"Scraping on %v goroutines every %s duration",
		concurrency,
		timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)

	// Run this loop every one minute
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)

		if err != nil {
			log.Println("Error fetching feeds: ", err)
			continue
		}

		waitGroup := &sync.WaitGroup{}

		for _, feed := range feeds {
			waitGroup.Add(1)

			// Spawning separate goRoutines
			go scrapeFeed(db, waitGroup, feed)
		}

		waitGroup.Wait()
	}
}

func scrapeFeed(
	db *database.Queries,
	waitGroup *sync.WaitGroup,
	feed database.Feed) {
	defer waitGroup.Done() // Decrements counter by 1

	_, err := db.MarkFeedFetched(context.Background(), feed.ID)

	if err != nil {
		log.Printf("Error marking feed as fetched: ", err)
		return
	}

	// rssFeed, err := urlToFeed(feed.Url)
	_, err = urlToFeed(feed.Url)

	if err != nil {
		log.Printf("Error feetching feed: ", err)
		return
	}

	// save to db
}
