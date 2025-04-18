package main

import (
	"context"
	"fmt"
	"time"
	"log"

	"github.com/google/uuid"
	"github.com/JustinPras/BlogAggregator/internal/database"
)

func handlerAggregator(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <duration>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Invalid duration entered: %w", err)
	}

	fmt.Printf("Collecting feeds every %v...\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("Error getting next feed to fetch: %w\n", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed %s as fetched: %w\n", feed.Name, err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Error fetching feed %s: %w", feed.Name, err)
	} 

	for _, item := range rssFeed.Channel.Item {
		parsedPubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing date: %w", err)
		}	

		createPostParams := database.CreatePostParams {
			ID:			uuid.New(),
			CreatedAt:	time.Now(),
			UpdatedAt: 	time.Now(),
			Title:		item.Title,
			Url:		item.Link,
			Description:item.Description,
			PublishedAt:parsedPubDate,
			FeedID:		feed.ID,
		}
		
		s.db.CreatePost(context.Background(), createPostParams)
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}