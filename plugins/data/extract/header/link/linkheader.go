// Package linkheader provides functions for parsing HTTP Link headers
package linkheader

import (
	"fmt"
	"regexp"
	"strings"
	// "sync"
)

// LINK_INDEX defines...
type LINK_INDEX string

const (

	// LINK_NEXT is the key for the next page
	LINK_NEXT LINK_INDEX = "next"

	// LINK_PREV is the key for the previous page
	LINK_PREV LINK_INDEX = "prev"

	// LINK_PREV is the key for the last page
	LINK_LAST LINK_INDEX = "last"

	//-- End
)

// linkRex sets...
var linkRex = regexp.MustCompile(`^.*<([^>]+)>; rel="next".*`)

// A Link is a single URL and related parameters
type Link struct {
	URL    string
	Rel    string
	Params map[string]string
}

// HasParam returns if a Link has a particular parameter or not
func (l Link) HasParam(key string) bool {
	for p := range l.Params {
		if p == key {
			return true
		}
	}
	return false
}

// Param returns the value of a parameter if it exists
func (l Link) Param(key string) string {
	for k, v := range l.Params {
		if key == k {
			return v
		}
	}
	return ""
}

// String returns the string representation of a link
func (l Link) String() string {
	p := make([]string, 0, len(l.Params))
	for k, v := range l.Params {
		p = append(p, fmt.Sprintf("%s=\"%s\"", k, v))
	}
	if l.Rel != "" {
		p = append(p, fmt.Sprintf("%s=\"%s\"", "rel", l.Rel))
	}
	return fmt.Sprintf("<%s>; %s", l.URL, strings.Join(p, "; "))
}

// Links is a slice of Link structs
type Links []Link

// FilterByRel filters a group of Links by the provided Rel attribute
func (l Links) FilterByRel(r string) Links {
	links := make(Links, 0)
	for _, link := range l {
		if link.Rel == r {
			links = append(links, link)
		}
	}
	return links
}

// ToString returns the string representation of multiple Links for use in HTTP responses etc
func (l Links) ToString() string {
	if l == nil {
		return fmt.Sprint(nil)
	}

	var strs []string
	for _, link := range l {
		strs = append(strs, link.String())
	}
	return strings.Join(strs, ", ")
}

func (l Links) getLink(key string) string {
	for _, link := range l {
		key := strings.ToLower(link.Rel)
		switch key {
		case "last":
			return link.String()
		case "next":
			return link.String()
		}
	}
	return ""
}

// String returns the string representation of the next link
func (l Links) Last() string {
	return l.getLink("last")
}

// String returns the string representation of the next link
func (l Links) Next() string {
	return l.getLink("next")
}

// ToMap returns the map[string]string representation of multiple links
func (l Links) ToMap() map[string]string {
	output := make(map[string]string, len(l))
	for _, link := range l {
		key := link.Rel
		output[key] = link.String()
	}
	return output
}

// Parse parses a raw Link header in the form:
//   <url>; rel="foo", <url>; rel="bar"; wat="dis"
// returning a slice of Link structs
func Parse(raw string) Links {
	var links Links

	// One chunk: <url>; rel="foo"
	for _, chunk := range strings.Split(raw, ",") {

		link := Link{URL: "", Rel: "", Params: make(map[string]string)}

		// Figure out what each piece of the chunk is
		for _, piece := range strings.Split(chunk, ";") {

			piece = strings.Trim(piece, " ")
			if piece == "" {
				continue
			}

			// URL
			if piece[0] == '<' && piece[len(piece)-1] == '>' {
				link.URL = strings.Trim(piece, "<>")
				continue
			}

			// Params
			key, val := parseParam(piece)
			if key == "" {
				continue
			}

			// Special case for rel
			if strings.ToLower(key) == "rel" {
				link.Rel = val
			} else {
				link.Params[key] = val
			}
		}

		if link.URL != "" {
			links = append(links, link)
		}

	}

	return links
}

// ParseMultiple is like Parse, but accepts a slice of headers
// rather than just one header string
func ParseMultiple(headers []string) Links {
	links := make(Links, 0)
	for _, header := range headers {
		links = append(links, Parse(header)...)
	}
	return links
}

// parseParam takes a raw param in the form key="val" and
// returns the key and value as seperate strings
func parseParam(raw string) (key, val string) {
	parts := strings.SplitN(raw, "=", 2)
	if len(parts) != 2 {
		return "", ""
	}
	key = parts[0]
	val = strings.Trim(parts[1], "\"")
	return key, val
}

// Example:
// Link: <https://api.github.com/organizations/3502508/repos?page=3>; rel="next",
// <https://api.github.com/organizations/3502508/repos?page=22>; rel="last",
// <https://api.github.com/organizations/3502508/repos?page=1>; rel="first",
// <https://api.github.com/organizations/3502508/repos?page=1>; rel="prev"
func parseLinkHeader(link string) string {
	submatch := linkRex.FindStringSubmatch(link)
	if len(submatch) != 2 {
		return ""
	}
	return submatch[1]
}
