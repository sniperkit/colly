package colly

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	// internal
	model "github.com/sniperkit/colly/pkg/model"
)

// elements represents a map of all available interface per element type
var elements = make(map[string]Extractor)

// Extractor represents a content extraction iterator
type Element interface {

	// Info represents...
	Info(ctx context.Context) (string, error)

	// Event represents...
	// Event(ctx context.Context, eventChan chan<- *model.EventResult, pattern string)

	//-- End
}

// registerElement function add a new service in the map/registry of new element types
func registerElement(element Element) {
	elements[Name(element)] = element
}

// Name returns the name of a service
func Name(element Element) string {
	parts := strings.Split(reflect.TypeOf(element).String(), ".")
	return strings.ToLower(parts[len(parts)-1])
}

// ForName returns the service for a given name, or an error if it doesn't exist
func ForName(name string) (Element, error) {
	if element, ok := extractors[strings.ToLower(name)]; ok {
		return element, nil
	}
	return &ElementNotFound{}, fmt.Errorf("service '%s' not found", name)
}

// NotFound is used when the specified service is not found
type ElementNotFound struct{}

// Info is not implemented
func (enf *ElementNotFound) Info(ctx context.Context) (string, error) {
	return "", errors.New("element info not found.")
}
