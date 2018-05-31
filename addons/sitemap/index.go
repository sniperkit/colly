package sitemap

import (
	"encoding/xml"
	"net/url"
	"strings"
)

func GetSitemapIndex(xmlSitemapURL url.URL) (SitemapIndex, error) {
	response, readErr := readURL(xmlSitemapURL)
	if readErr != nil {
		return SitemapIndex{}, readErr
	}
	if !strings.Contains(string(response.GetBody()), "</sitemapindex>") {
		return SitemapIndex{}, SitemapIndexError{"Invalid content"}
	}
	var sitemapIndex SitemapIndex
	unmarshalError := xml.Unmarshal(response.GetBody(), &sitemapIndex)
	if unmarshalError != nil {
		return SitemapIndex{}, unmarshalError
	}
	return sitemapIndex, nil
}

func (sitemapIndexError SitemapIndexError) Error() string {
	return sitemapIndexError.message
}

func isInvalidSitemapIndexContent(err error) bool {
	if err == nil {
		return false
	}

	return err.Error() == "Invalid content"
}
