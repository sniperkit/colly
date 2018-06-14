package colly

import (
	linkheader "github.com/sniperkit/colly/plugins/data/extract/header/link"
)

/*
	Input example:
	- `<https://api.github.com/user/58276/repos?page=2>; rel="next", <https://api.github.com/user/58276/repos?page=2>; rel="last"`
*/

func (c *Collector) linkHeaderFromString(header, reqUrl string, reqId uint32) map[string]string {
	links := linkheader.Parse(header)
	output := make(map[string]string, len(links))
	for _, link := range links {
		if c.debugger != nil {
			c.debugger.Event(createEvent("header-link", reqId, c.ID, map[string]string{
				"link": link.String(),
				"url":  reqUrl,
			}))
		}
		output[link.Rel] = link.URL
	}
	return output
}

func (c *Collector) linkHeaderFromSlice(headers []string, reqUrl string, reqId uint32) map[string]string {
	links := linkheader.ParseMultiple(headers)
	output := make(map[string]string, len(links))
	for _, link := range links {
		if c.debugger != nil {
			c.debugger.Event(createEvent("header-link", reqId, c.ID, map[string]string{
				"link": link.String(),
				"url":  reqUrl,
			}))
		}
		output[link.Rel] = link.URL
	}
	return output
}

func (c *Collector) linkHeaderFilterByRel(header, reqUrl string, reqId uint32) map[string]string {
	links := linkheader.Parse(header)
	output := make(map[string]string, len(links))
	for _, link := range links.FilterByRel("last") {
		if c.debugger != nil {
			c.debugger.Event(createEvent("header-link", reqId, c.ID, map[string]string{
				"link": link.String(),
				"url":  reqUrl,
			}))
		}
		output[link.Rel] = link.URL
	}
	return output
}

func (c *Collector) linkHeaderString(link *linkheader.Link, reqUrl string, reqId uint32) string {
	if c.debugger != nil {
		c.debugger.Event(createEvent("header-link", reqId, c.ID, map[string]string{
			"link": link.String(),
			"url":  reqUrl,
		}))
	}
	return link.String()
}

/*
func (c *Collector) linkHeaderJoin(links *linkheader.Links, reqUrl string, reqId uint32) string {
	if c.debugger != nil {
		c.debugger.Event(createEvent("header-link", reqId, c.ID, map[string]string{
			"links": links.String(),
			"url":   reqUrl,
		}))
	}
	return links.String()
}
*/
