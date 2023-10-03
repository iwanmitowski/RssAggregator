package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/iwanmitowski/RssAggregator/internal/database"
)

func startScraping(
	db database.Database,
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
	db database.Database,
	waitGroup *sync.WaitGroup,
	feed database.Feed) {
	defer waitGroup.Done() // Decrements counter by 1

	_, err := db.MarkFeedFetched(context.Background(), feed.ID)

	if err != nil {
		log.Printf("Error marking feed as fetched: ", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Printf("Error feetching feed: ", err)
		return
	}

	// save to db
	for _, item := range rssFeed.Channel.Item {
		desciption := sql.NullString{}

		if item.Description != "" {
			desciption.String = item.Description
			desciption.Valid = true
		}

		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil {
			log.Printf("couldn't parse date %v with err: ", item.PubDate, err)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: desciption,
			PublishedAt: publishedAt,
			Url:         item.Link,
			FeedID:      feed.ID,
		})

		if err != nil {
			// Fix this
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("failed to create post: ", err)
		}
	}

	log.Printf("Feeds %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
