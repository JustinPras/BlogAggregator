package main

import (
	"context"
	"fmt"
	"time"
	"log"
	"strconv"

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

	layoutTime := "Mon, 02 Jan 2006 15:04:05 -0700"
	 

	for _, item := range rssFeed.Channel.Item {
		parsedPubDate, err := time.Parse(layoutTime, item.PubDate)
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

func handlerBrowse(s *state, cmd command, currentUser database.User) error {
	var limit int32
	limit = 2
	if len(cmd.Args) == 1 {
		num, err := strconv.ParseInt(cmd.Args[0], 10, 0)
		if err != nil {
			return fmt.Errorf("Error parsing specified limit %s: %w", cmd.Args[0], err)
		}
		limit = int32(num)
	}

	

	getPostsByUserParams := database.GetPostsByUserParams{
		UserID:	currentUser.ID,
		Limit:	limit,
	}

	posts, err := s.db.GetPostsByUser(context.Background(), getPostsByUserParams)
	if err != nil {
		return fmt.Errorf("Error retrieving posts for user %s: %w", currentUser.Name, err)
	}

	for _, post := range(posts) {
		printPost(post)
		fmt.Println("=====================================")
	}

	return nil
}

func printPost(post database.Post) {
	fmt.Printf("* ID:            %s\n", post.ID)
	fmt.Printf("* Created:       %v\n", post.CreatedAt)
	fmt.Printf("* Updated:       %v\n", post.UpdatedAt)
	fmt.Printf("* Title:         %s\n", post.Title)
	fmt.Printf("* URL:           %s\n", post.Url)
	fmt.Printf("* Description:   %s\n", post.Description)
	fmt.Printf("* Published:     %v\n", post.PublishedAt)
	fmt.Printf("* Feed ID:       %s\n", post.FeedID)
}