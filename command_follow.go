package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/FT1006/blogator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		fmt.Println("Error: Not enough arguments provided")
		os.Exit(1) // Exit with non-zero status code to indicate an error
	}

	url := cmd.Args[0]

	targetFeed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		fmt.Printf("error getting feed: %v", err)
		return nil
	}

	newFeedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    targetFeed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		fmt.Printf("error creating feed follow: %v", err)
		return nil
	}

	fmt.Println(newFeedFollow)
	return nil

}
