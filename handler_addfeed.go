package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/JustinPras/BlogAggregator/internal/database"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	currentUserName := s.cfg.CurrentUserName
	currentUser, err := s.db.GetUser(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("error retrieving current user: %w", err)
	}	
	userID := currentUser.ID

	name := cmd.Args[0]
	url := cmd.Args[1]
	
	feedParams := database.CreateFeedParams {
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name:		name,
		Url:		url,
		UserID:		userID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("couldn't create new feed: %w", err)
	}

	fmt.Println("Feed created successfully!")
	fmt.Printf("Feed Data: %+v\n", feed)
	return nil
}