package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/lululululu5/blog-aggregator/internal/database"
)


func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Collecting feeds every %s on %v goroutines...", timeBetweenRequest, concurrency)
	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Couldn't get next feeds to fetch", err)
			continue
		}
		log.Printf("Found %v feeds to fetch!", len(feeds))
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Items {
		publishDate, err := convStringToTime(item.PubDate)
		if err != nil {
			log.Printf("Couldn't decode publishing date. We will use current date")
			publishDate = time.Now().UTC()
		}

		post, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: item.Title,
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid: true,
			},
			PublishedAt: publishDate,
			FeedID: feed.ID,
		})
		if err != nil {
			log.Println("Could not create new post")
			continue
		}
		
		log.Println("Successfully added post: ", post)

	}
}


func fetchFeed(url string) (*RssFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RssFeed
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return nil, errors.New("could not decode RSS feed")
	}
	
	return &rssFeed, nil
}

func (cfg *apiConfig) GetNextFeedsToFetch(n int32) ([]Feed, error){
	ctx := context.Background()
	feeds, err := cfg.DB.GetNextFeedsToFetch(ctx, n)
	if err != nil {
		return nil, errors.New("could not fetch feeds")
	}

	result := databaseFeedsToFeeds(feeds)
	return result, nil
}


func convStringToTime(dateString string) (time.Time, error) {
	layout := "Mon, 02 Jan 2006 15:04:05 -0700"
	t, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}