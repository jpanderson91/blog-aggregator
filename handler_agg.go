package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jpanderson91/blog-aggregator/internal/database"
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

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("couldn't get next feed to fetch:", err)
		return
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("couldn't mark feed as fetched:", err)
		return
	}

	fetched, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Println("couldn't fetch feed:", err)
		return
	}

	for _, item := range fetched.Channel.Item {
		publishedAt := sql.NullTime{}
		if item.PubDate != "" {
			t, err := time.Parse(time.RFC1123Z, item.PubDate)
			if err == nil {
				publishedAt = sql.NullTime{Time: t, Valid: true}
			}
		}

		description := sql.NullString{}
		if item.Description != "" {
			description = sql.NullString{String: item.Description, Valid: true}
		}

		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if isDuplicateKeyError(err) {
				continue
			}
			log.Println(err)
			continue
		}
	}
}