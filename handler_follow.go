package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jpanderson91/blog-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: follow <url>")
	}
	url := cmd.Args[0]

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	now := time.Now()
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Printf("Feed: %s\n", feedFollow.FeedName)
	fmt.Printf("User: %s\n", feedFollow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}
	return nil
}
