package main

import "encoding/xml"

// type RssFeed struct {
// 	XMLName xml.Name `xml:"rss"`
// 	Text    string   `xml:",chardata"`
// 	Version string   `xml:"version,attr"`
// 	Channel struct {
// 		Text  string `xml:",chardata"`
// 		Title string `xml:"title"`
// 		Link  struct {
// 			Text string `xml:",chardata"`
// 			Href string `xml:"href,attr"`
// 			Rel  string `xml:"rel,attr"`
// 			Type string `xml:"type,attr"`
// 		} `xml:"link"`
// 		Description   string `xml:"description"`
// 		Generator     string `xml:"generator"`
// 		Language      string `xml:"language"`
// 		LastBuildDate string `xml:"lastBuildDate"`
// 		Item          []struct {
// 			Text  string `xml:",chardata"`
// 			Title struct {
// 				Text   string `xml:",chardata"`
// 				Insert struct {
// 					Text       string `xml:",chardata"`
// 					Technology string `xml:"technology,attr"`
// 				} `xml:"insert"`
// 			} `xml:"title"`
// 			Link        string `xml:"link"`
// 			PubDate     string `xml:"pubDate"`
// 			Guid        string `xml:"guid"`
// 			Description string `xml:"description"`
// 		} `xml:"item"`
// 	} `xml:"channel"`
// }

type RssFeed struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title         string   `xml:"title"`
	Link          Link     `xml:"link"`
	Description   string   `xml:"description"`
	Generator     string   `xml:"generator"`
	Language      string   `xml:"language"`
	LastBuildDate string   `xml:"lastBuildDate"`
	Items         []Item   `xml:"item"`
}

type Link struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type Item struct {
	Title       Title `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
	Description string `xml:"description"`
}

type Title struct {
	Text   string  `xml:",chardata"`
	Insert *Insert `xml:"insert,omitempty"`
}

type Insert struct {
	Text       string `xml:",chardata"`
	Technology string `xml:"technology,attr"`
}