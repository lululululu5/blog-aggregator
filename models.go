package main

import (
	"database/sql"
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

type PostDB struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt time.Time
	FeedID      uuid.UUID
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	var description string
	if post.Description.Valid {
		description = post.Description.String
	} else {
		description = ""
	}
	return Post{
		ID: post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Title: post.Title,
		Url: post.Url,
		Description: description,
		PublishedAt: post.PublishedAt,
		FeedID: post.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = databasePostToPost(post)
	}
	return result
}
