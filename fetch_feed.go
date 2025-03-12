package main

import (
	"net/http"
	"io"
	"time"
	"encoding/xml"
	"context"
	"html"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Set("user-agent", "gator")

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	rssFeedResp := RSSFeed{}
	err = xml.Unmarshal(data, &rssFeedResp)
	if err != nil {
		return &RSSFeed{}, err
	}

	rssFeedResp.Channel.Title = html.UnescapeString(rssFeedResp.Channel.Title)
	rssFeedResp.Channel.Description = html.UnescapeString(rssFeedResp.Channel.Description)
	for _, item := range rssFeedResp.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}


	return &rssFeedResp, nil
}