package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feedFetched, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	cleanFeed(feedFetched)
	fmt.Printf("%+v\n", feedFetched)
	return nil
}
