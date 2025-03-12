package main

import (
	"context"
	"fmt"
)

func handlerAggregator(s *state, cmd command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("RSS Feed: %+v\n", rssFeed)

	return nil
}