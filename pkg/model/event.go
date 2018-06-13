package model

import (
	"errors"
	"time"

	"github.com/skratchdot/open-golang/open"
)

const (
	defaultWho   = "somebody"
	defaultWhat  = "did something with"
	defaultWhich = "some item"
)

var eventTypes = map[string]string{
	"CreateEvent":  "created",
	"DeleteEvent":  "deleted",
	"MemberEvent":  "added someone to",
	"PublicEvent":  "made public",
	"ReleaseEvent": "released",
	"WatchEvent":   "starred",
}

// Event is a git-hosting service event
type Event struct {
	Who   string
	What  string
	Which string
	URL   string
	When  time.Time
}

// EventResult wraps an event and an error
type EventResult struct {
	Event *Event
	Error error
}

// OpenInBrowser opens the event in the browser
func (event *Event) OpenInBrowser() error {
	if event.URL == "" {
		return errors.New("no URL for event")
	}
	return open.Start(event.URL)
}
