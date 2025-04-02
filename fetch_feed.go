package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/FT1006/blogator/internal/database"
	"github.com/google/uuid"
)

func newClient() *http.Client {
	return &http.Client{}
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	var rssFeed RSSFeed
	client := newClient()
	req, err := http.NewRequestWithContext(ctx, "", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "blogator")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return nil, err
	}
	return &rssFeed, nil

}

func scrapeFeeds(ctx context.Context, s *state) {
	toFetch, err := s.db.GetAllFeeds(ctx)
	if err != nil {
		fmt.Printf("Error fetchig feeds: %v\n", err)
	}

	noToFetch := len(toFetch)

	for i := 0; i < noToFetch; i++ {
		feed, err := s.db.GetNextFeedToFetch(ctx)
		if err != nil {
			fmt.Printf("Error getting next feed to fetch: %v\n", err)
		}
		fetchedFeed, err := fetchFeed(ctx, feed.Url)
		if err != nil {
			fmt.Printf("Error fetchig next feed: %v\n", err)
		}

		// Mark the feed as fetched
		err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
			LastFetchedAt: sql.NullTime{
				Time: time.Now(),
			},
			UpdatedAt: time.Now(),
			ID:        feed.ID,
		})
		if err != nil {
			fmt.Printf("Error marking feed as fetched: %v\n", err)
		} else {
			if err != nil {
				fmt.Printf("Error parsing pubDate: %v\n", err)
			}
			fetchedFeed = cleanFeed(fetchedFeed)
			for _, item := range fetchedFeed.Channel.Item {
				publishedAt, err := time.Parse(time.RFC1123, item.PubDate)
				if err != nil {
					fmt.Printf("Error parsing pubDate: %v\n", err)
				}
				if _, err = s.db.CreatePost(ctx, database.CreatePostParams{
					ID:          uuid.New(),
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					Title:       item.Title,
					Url:         item.Link,
					Description: sql.NullString{String: item.Description, Valid: true},
					PublishedAt: publishedAt,
					FeedID:      feed.ID,
				}); err != nil {
					fmt.Printf("Error creating post: %v\n", err)
				}
			}
		}
	}
}
