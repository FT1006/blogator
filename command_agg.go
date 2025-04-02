package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) > 1 {
		fmt.Println("Error: too many arguments provided")
		os.Exit(1) // Exit with non-zero status code to indicate an error
	}

	var timeDurationString string
	if len(cmd.Args) == 1 {
		timeDurationString = cmd.Args[0]
	} else {
		timeDurationString = "30m"
	}

	timeBetweenRequests, err := time.ParseDuration(timeDurationString)
	if err != nil {
		fmt.Printf("Error: Invalid time duration provided: %v\n", err)
		os.Exit(1) // Exit with non-zero status code to indicate an error
	}

	fmt.Printf("Collecting feeds every %s\n", timeDurationString)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		fmt.Println("Fetching feeds...")
		scrapeFeeds(context.Background(), s)
	}
}
