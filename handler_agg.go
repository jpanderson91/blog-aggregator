package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
    return fmt.Errorf("usage: agg <time_between_reqs>")
	}
	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't parse duration: %w", err)
	}
	fmt.Printf("Collecting feeds every %s\n", duration)
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
    scrapeFeeds(s)
	}

}
func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
if err != nil {
    return err
}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	fetched, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	for _, item := range fetched.Channel.Item {
    fmt.Println(item.Title)
}
return nil
}