package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		fmt.Println("Error: No time duration provided")
		os.Exit(1) // Exit with non-zero status code to indicate an error
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		fmt.Printf("Error: Invalid time duration provided: %v\n", err)
		os.Exit(1) // Exit with non-zero status code to indicate an error
	}

	fmt.Printf("Collecting feeds every %s\n", cmd.Args[0])

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		fmt.Println("Fetching feeds...")
		scrapeFeeds(context.Background(), s)
		if err != nil {
			return err
		}
	}
}
