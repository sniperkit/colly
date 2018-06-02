package sitemap

import (
	"fmt"
	"net/url"
)

func checkSitemap(loc string) bool {
	return true
}

func GetURLs(xmlSitemapURL url.URL) ([]url.URL, error) {
	var urls []url.URL

	urlsFromIndex, indexError := GetURLsFromSitemapIndex(xmlSitemapURL)
	if indexError == nil {
		urls = urlsFromIndex
	}

	urlsFromSitemap, sitemapError := GetURLsFromSitemap(xmlSitemapURL)
	if sitemapError == nil {
		urls = append(urls, urlsFromSitemap...)
	}

	if isInvalidSitemapIndexContent(indexError) && isInvalidXMLSitemapContent(sitemapError) {
		return nil, fmt.Errorf("%q is neither a sitemap index nor a XML sitemap", xmlSitemapURL.String())
	}

	return urls, nil
}

func GetURLsFromSitemap(xmlSitemapURL url.URL) ([]url.URL, error) {
	var urls []url.URL

	sitemap, xmlSitemapError := GetXMLSitemap(xmlSitemapURL)
	if xmlSitemapError != nil {
		return nil, xmlSitemapError
	}

	for _, urlEntry := range sitemap.URLs {
		parsedURL, parseError := url.Parse(urlEntry.Location)
		if parseError != nil {
			return nil, parseError
		}
		urls = append(urls, *parsedURL)
	}

	return urls, nil
}

func GetURLsFromSitemapIndex(xmlSitemapURL url.URL) ([]url.URL, error) {
	var urls []url.URL

	sitemapIndex, sitemapIndexError := GetSitemapIndex(xmlSitemapURL)
	if sitemapIndexError != nil {
		return nil, sitemapIndexError
	}

	for _, sitemap := range sitemapIndex.Sitemaps {
		locationURL, err := url.Parse(sitemap.Location)
		if err != nil {
			return nil, err
		}
		sitemapUrls, err := GetURLsFromSitemap(*locationURL)
		if err != nil {
			return nil, err
		}
		urls = append(urls, sitemapUrls...)
	}
	return urls, nil

}
