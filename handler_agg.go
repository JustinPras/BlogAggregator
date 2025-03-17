package main

import (
	"context"
	"fmt"
	"time"
	"database/sql"

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

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting next feed to fetch: %w", err)
	}

	lastFetchedAt := sql.NullTime{Time: time.Now(), Valid: true}

	markFeedFetchedParams := database.MarkFeedFetchedParams {
		LastFetchedAt:	lastFetchedAt,
		UpdatedAt:		time.Now(),
		ID:				feed.ID,
	}

	err = s.db.MarkFeedFetched(context.Background(), markFeedFetchedParams)
	if err != nil {
		return fmt.Errorf("Error marking feed as fetched: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %w", err)
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("* Title:     %s\n", item.Title)
	}
	return nil
}