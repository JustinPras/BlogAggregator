package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/JustinPras/BlogAggregator/internal/database"
)

func handlerDisplayFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving feeds: %w", err)
	}

	if len(feeds) == 0 {
		return fmt.Errorf("No feeds found.")
	}
	
	for _, feed := range(feeds) {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't retrieve feed creator: %w", err)
		}

		printFeed(feed, user)
		fmt.Println("=====================================")
	}
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error retrieving current user: %w", err)
	}	

	name := cmd.Args[0]
	url := cmd.Args[1]
	
	feedParams := database.CreateFeedParams {
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name:		name,
		Url:		url,
		UserID:		currentUser.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("couldn't create new feed: %w", err)
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

	fmt.Println("Feed created successfully!")
	printFeed(feed, currentUser)
	fmt.Println("=====================================")
	fmt.Printf("User '%s' followed the feed '%s' successfully!\n", feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("=====================================")
	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
}