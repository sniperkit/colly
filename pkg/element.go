package colly

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	// internal
	model "github.com/sniperkit/colly/pkg/model"
)

// extractors represents a map of all available interface per service
var extractors = make(map[string]Extractor)

// Extractor represents a content extraction iterator
type Extractor interface {

	// Info represents...
	Info(ctx context.Context) (string, error)

	// Event represents...
	// Event(ctx context.Context, eventChan chan<- *model.EventResult, pattern string)

	//-- End
}

// registerService function add a new service in the map/registry of services
func registerExtractor(extractor Extractor) {
	extractors[Name(extractor)] = extractor
}

// Name returns the name of a service
func Name(extractor Extractor) string {
	parts := strings.Split(reflect.TypeOf(service).String(), ".")
	return strings.ToLower(parts[len(parts)-1])
}

// ForName returns the service for a given name, or an error if it doesn't exist
func ForName(name string) (Service, error) {
	if service, ok := extractors[strings.ToLower(name)]; ok {
		return service, nil
	}
	return &ExtractorNotFound{}, fmt.Errorf("service '%s' not found", name)
}

// NotFound is used when the specified service is not found
type ExtractorNotFound struct{}

// Info is not implemented
func (enf *ExtractorNotFound) Info(ctx context.Context) (string, error) {
	return "", errors.New("service not found")
}
