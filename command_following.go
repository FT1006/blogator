package main

import (
	"context"
	"fmt"

	"github.com/FT1006/blogator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		fmt.Printf("error getting feed follows: %v", err)
		return nil
	}

	for _, feedFollow := range feedFollows {
		fmt.Println(feedFollow.FeedsName)
	}

	return nil
}
