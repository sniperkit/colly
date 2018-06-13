package service

import (
	"context"
	"errors"

	// internal
	model "github.com/sniperkit/colly/pkg/model"
)

// NotFound is used when the specified service is not found
type NotFound struct {
}

// Login is not implemented
func (nf *NotFound) Login(ctx context.Context) (string, error) {
	return "", errors.New("service not found")
}

// GetEvents is not implemented
func (nf *NotFound) GetEvents(ctx context.Context, eventChan chan<- *model.EventResult, token, user string, page, count int) {
}
