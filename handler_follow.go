package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/JustinPras/BlogAggregator/internal/database"
)

func handlerFollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
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

func handlerListFeedFollows(s *state, cmd command, currentUser database.User) error {

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

func handlerUnfollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Feed does not exist: %w", err)
	}

	userUrlParams := database.DeleteFeedFollowByUserAndURLParams {
		UserID:	currentUser.ID,
		FeedID:	feed.ID,	
	}

	err = s.db.DeleteFeedFollowByUserAndURL(context.Background(), userUrlParams)
	if err != nil {
		return fmt.Errorf("Error unfollowing feed: %w", err)
	}

	fmt.Printf("%s unfollowed successfully!\n", feed.Name)
	return nil
}