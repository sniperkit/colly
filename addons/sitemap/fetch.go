package sitemap

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/sniperkit/colly/pkg"
)

func readURL(url url.URL) (colly.Response, error) {
	startTime := time.Now().UTC()
	resp, fetchErr := http.Get(url.String())
	if fetchErr != nil {
		return colly.Response{}, fetchErr
	}

	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return colly.Response{}, readErr
	}

	endTime := time.Now().UTC()

	// content type
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = http.DetectContentType(body)
	}

	return colly.Response{
		Body:        body,
		StatusCode:  resp.StatusCode,
		StartTime:   startTime,
		EndTime:     endTime,
		ContentType: contentType,
	}, nil
}
