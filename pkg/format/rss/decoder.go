package rss

import (
	"io"

	// external
	gofeed "github.com/sniperkit/gofeed/pkg"

	// internal
	colly "github.com/sniperkit/colly/pkg"
	decoder "github.com/sniperkit/colly/pkg/decoder"
)

// RSS is the key for the rss encoding
const RSS = "rss"

// ContentTypeRSS...
type RSSDecoder struct {
	resp *colly.Response
	err  error
}

// Check...
func (c *RSSDecoder) Check(resp *colly.Response) error { return decoder.ErrNotImplementedYet }

// Decode...
func (c *RSSDecoder) Decode(resp *colly.Response, v *map[string]interface{}) error {
	return decoder.ErrNotImplementedYet
}

// NewRSSDecoder returns the RSS decoder
func NewRSSDecoder(isCollection bool) Decoder {
	if isCollection {
		return RSSCollectionDecoder
	}
	return RSSDecoder
}

// RSSDecoder implements the Decoder interface
func RSSDecoder(r io.Reader, v *map[string]interface{}) error {
	fp := gofeed.NewParser()
	feed, err := fp.Parse(r)
	if err != nil {
		return err
	}
	*(v) = map[string]interface{}{
		"items":       feed.Items,
		"author":      feed.Author,
		"categories":  feed.Categories,
		"custom":      feed.Custom,
		"copyright":   feed.Copyright,
		"description": feed.Description,
		"type":        feed.FeedType,
		"language":    feed.Language,
		"title":       feed.Title,
		"published":   feed.Published,
		"updated":     feed.Updated,
	}
	if feed.Image != nil {
		(*v)["img_url"] = feed.Image.URL
	}
	return nil
}

// RSSCollectionDecoder implements the Decoder interface over a collection
func RSSCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	fp := gofeed.NewParser()
	feed, err := fp.Parse(r)
	if err != nil {
		return err
	}
	f := map[string]interface{}{
		"items":       feed.Items,
		"author":      feed.Author,
		"categories":  feed.Categories,
		"custom":      feed.Custom,
		"copyright":   feed.Copyright,
		"description": feed.Description,
		"type":        feed.FeedType,
		"language":    feed.Language,
		"title":       feed.Title,
		"published":   feed.Published,
		"updated":     feed.Updated,
	}
	if feed.Image != nil {
		(*v)["img_url"] = feed.Image.URL
	}
	*(v) = map[string]interface{}{"collection": f}
	return nil
}
