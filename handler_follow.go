package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/JustinPras/BlogAggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error retrieving current user: %w", err)
	}	

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("That feed does not exist: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams {
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		UserID:		currentUser.ID,
		FeedID:		feed.ID,
	}


	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Error following feed: %w", err)
	}

	fmt.Printf("User '%s' followed the feed '%s' successfully!\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerListFeedFollows(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error retrieving current user: %w", err)
	}	

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("Error retrieving followed feeds for current user: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Println("Current user is following the listed feeds:")
	for _, feedFollow := range(feedFollows) {
		fmt.Printf("* %s\n", feedFollow.FeedName)
	}
	return nil
}