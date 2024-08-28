package main

type RssFeed struct {
	Channel struct {
		Title         string   `xml:"title"`
		Link          string    `xml:"link"`
		Description   string   `xml:"description"`
		Generator     string   `xml:"generator"`
		Language      string   `xml:"language"`
		LastBuildDate string   `xml:"lastBuildDate"`
		Items         []RSSItem   `xml:"item"`
	}  `xml:"channel"`
}


type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
	Description string `xml:"description"`
}
