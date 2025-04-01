package main

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
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
