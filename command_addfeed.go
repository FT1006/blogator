package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/FT1006/blogator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		fmt.Println("no feed name or url")
		os.Exit(1) // Exit with non-zero status code to indicate an error
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	newFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	addedFollowedFeeds, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    newFeed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		fmt.Printf("error creating feed follow: %v", err)
		return nil
	}

	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		fmt.Printf("error getting followed feeds: %v", err)
		return nil
	}

	fmt.Printf("No. of followed feeds: %d\n", len(followedFeeds))

	fmt.Printf("Feed %s created successfully\n", newFeed.Name)
	fmt.Printf("ID: %s\n", newFeed.ID)
	fmt.Printf("Created at: %s\n", newFeed.CreatedAt)
	fmt.Printf("Updated at: %s\n", newFeed.UpdatedAt)
	fmt.Printf("URL: %s\n", newFeed.Url)

	fmt.Printf("Followed feeds updated successfully\n")
	for _, followedFeed := range addedFollowedFeeds {
		fmt.Printf("Followed feed: %s\n", followedFeed.FeedsName)
	}

	return nil
}
