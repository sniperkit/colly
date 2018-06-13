package storage

import (
	"context"
	"errors"
	"net/http/cookiejar"
	"net/url"
	"sync"
)

// InMemoryStorage is the default storage backend of colly.
// InMemoryStorage keeps cookies and visited urls in memory
// without persisting data on the disk.
type InMemoryStorage struct {
	visitedURLs map[uint64]bool
	lock        *sync.RWMutex
	jar         *cookiejar.Jar
}

// Init initializes InMemoryStorage
func (s *InMemoryStorage) Init() error {
	if s.visitedURLs == nil {
		s.visitedURLs = make(map[uint64]bool)
	}
	if s.lock == nil {
		s.lock = &sync.RWMutex{}
	}
	if s.jar == nil {
		var err error
		s.jar, err = cookiejar.New(nil)
		return err
	}
	return nil
}

// WithConfig is not implemented
func (s *InMemoryStorage) WithConfig(ctx context.Context) error {
	return errors.New("implemented ready yet")
}

// Info is not implemented
func (s *InMemoryStorage) Info(ctx context.Context) (string, error) {
	return "", errors.New("not implemented yet")
}

// Get is not implemented
func (s *InMemoryStorage) Get(ctx context.Context) (string, error) {
	return "", errors.New("not implemented yet")
}

// Set is not implemented
func (s *InMemoryStorage) Set(ctx context.Context) (string, error) {
	return "", errors.New("not implemented yet")
}

// Update is not implemented
func (s *InMemoryStorage) Update(ctx context.Context) (string, error) {
	return "", errors.New("not implemented yet")
}

// Delete is not implemented
func (s *InMemoryStorage) Delete(ctx context.Context) error { return errors.New("storage not found") }

// Health is not implemented
func (s *InMemoryStorage) Health(ctx context.Context) (map[string]bool, error) {
	return nil, errors.New("storage not found")
}

// Close is not implemented
func (s *InMemoryStorage) Close(ctx context.Context) error { return nil }

// Visited implements Storage.Visited()
func (s *InMemoryStorage) visited(requestID uint64) error {
	s.lock.Lock()
	s.visitedURLs[requestID] = true
	s.lock.Unlock()
	return nil
}

// IsVisited implements Storage.IsVisited()
func (s *InMemoryStorage) isVisited(requestID uint64) (bool, error) {
	s.lock.RLock()
	visited := s.visitedURLs[requestID]
	s.lock.RUnlock()
	return visited, nil
}

// Cookies implements Storage.Cookies()
func (s *InMemoryStorage) cookies(u *url.URL) string {
	return StringifyCookies(s.jar.Cookies(u))
}

// SetCookies implements Storage.SetCookies()
func (s *InMemoryStorage) setCookies(u *url.URL, cookies string) {
	s.jar.SetCookies(u, UnstringifyCookies(cookies))
}
