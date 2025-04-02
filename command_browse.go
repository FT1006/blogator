package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error {

	if len(cmd.Args) > 1 {
		fmt.Println("Error: too many arguments provided")
		os.Exit(1) // Exit with non-zero status code to indicate an error
	}

	limit := 2
	if len(cmd.Args) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			fmt.Printf("Error converting limit to integer: %v\n", err)
			os.Exit(1) // Exit with non-zero status code to indicate an error
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), int32(limit))
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("Published at: %s\n", post.PublishedAt)
		fmt.Printf("Fetched at: %s\n", post.CreatedAt)
		fmt.Println()
	}
	return nil
}
