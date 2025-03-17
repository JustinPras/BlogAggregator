package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JustinPras/BlogAggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, currentUser database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		num, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("Invalid limit: %w", err)
		}
		limit = num
	}

	getPostsByUserParams := database.GetPostsByUserParams{
		UserID:	currentUser.ID,
		Limit:	int32(limit),
	}

	posts, err := s.db.GetPostsByUser(context.Background(), getPostsByUserParams)
	if err != nil {
		return fmt.Errorf("Error retrieving posts for user %s: %w", currentUser.Name, err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), currentUser.Name)
	for _, post := range(posts) {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}