package colly

import (
	"io"

	gofeed "github.com/sniperkit/gofeed/pkg"
)

// RSS is the key for the rss encoding
const RSS = "rss"

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

// XMLCollectionDecoder implements the Decoder interface over a collection
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

// ContentTypeRSS...
type ContentTypeRSS struct {
	resp *Response
	err  error
}

// Check...
func (c *ContentTypeRSS) Check(resp *Response) error { return ErrNotImplementedYet }
