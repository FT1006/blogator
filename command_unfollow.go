package main

import (
	"context"
	"fmt"
	"os"

	"github.com/FT1006/blogator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
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

	err = s.db.DeleteFeedFollowsForUser(
		context.Background(),
		database.DeleteFeedFollowsForUserParams{
			UserID: user.ID,
			FeedID: targetFeed.ID,
		},
	)

	if err != nil {
		fmt.Printf("error deleting feed follow: %v", err)
		return nil
	}
	return nil
}
