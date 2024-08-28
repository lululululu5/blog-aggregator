package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/lululululu5/blog-aggregator/internal/database"
)

type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	ApiKey string `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name: user.Name,
		ApiKey: user.ApiKey,
	}	
}

type Feed struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	Url string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	var lastFechtedAt *time.Time
	if feed.LastFetchedAt.Valid {
		lastFechtedAt = &feed.LastFetchedAt.Time
	} else {
		lastFechtedAt = nil
	}

	return Feed{
		ID: feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name: feed.Name,
		Url: feed.Url,
		UserID: feed.UserID,
		LastFetchedAt: lastFechtedAt,
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	result := make([]Feed, len(feeds))
	for i, feed := range feeds {
		result[i] = databaseFeedToFeed(feed)
	}
	return result
}

type FeedFollow struct {
	ID uuid.UUID `json:"id"`
	FeedID uuid.UUID `json:"feed_id"`
	UserID uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedFollowToFeedFoolow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID: feedFollow.ID,
		FeedID: feedFollow.FeedID,
		UserID: feedFollow.UserID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
	}
}

func databaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	result := make([]FeedFollow, len(feedFollows))
	for i, feedFollow := range feedFollows {
		result[i] = databaseFeedFollowToFeedFoolow(feedFollow)
	}
	return result
}
