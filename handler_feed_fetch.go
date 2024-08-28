package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lululululu5/blog-aggregator/internal/database"
)




func (cfg *apiConfig) FetchFeedData(url string) ([]RssFeedResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("could not connect to URL")
	}
	defer resp.Body.Close()
	
	decoderXML := xml.NewDecoder(resp.Body)
	rssFeed := RssFeed{}
	err = decoderXML.Decode(&rssFeed)
	if err != nil {
		return nil, errors.New("could not decode RSS feed")
	}
	
	result := itemsToRssFeeds(rssFeed.Channel.Items)
	return result, nil
}

func (cfg *apiConfig) GetNextFeedsToFetch(n int32) ([]Feed, error){
	ctx := context.TODO()
	feeds, err := cfg.DB.GetNextFeedsToFetch(ctx, n)
	if err != nil {
		return nil, errors.New("could not fetch feeds")
	}

	result := databaseFeedsToFeeds(feeds)
	return result, nil
}


func (cfg *apiConfig) MarkFeedFetched(feedID uuid.UUID) error {
	ctx := context.TODO()
	err := cfg.DB.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		ID: feedID,
		UpdatedAt: time.Now().UTC(),
		LastFetchedAt: sql.NullTime{
			Time: time.Now().UTC(),
			Valid: true,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (cfg *apiConfig) workerFeed(numFeeds int32) {
	feedsToFetch, err := cfg.GetNextFeedsToFetch(numFeeds)
	if err != nil {
		log.Println("Something went terribly wrong")
	}

	for _, feed := range feedsToFetch {
		log.Printf("Fetching the following feed: %v\n", feed.Url)
		_, err := cfg.FetchFeedData(feed.Url)
		if err != nil {
			log.Println("Could not fetch data")
		} else {
			log.Println("Feeds successfull fetched")
			err = cfg.MarkFeedFetched(feed.ID)
			if err != nil {
				log.Println("Could not mark feed as fetched")
			}
		}
		
	}

}