package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		fmt.Printf("error getting all feeds: %v", err)
		return nil
	}

	for _, feed := range feeds {
		creator, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			fmt.Printf("error getting user: %v", err)
			return nil
		}
		fmt.Printf("* %s\n", feed.Name)
		fmt.Printf("  URL: %s\n", feed.Url)
		fmt.Printf("  Created by: %s\n", creator.Name)
	}
	return nil
}
