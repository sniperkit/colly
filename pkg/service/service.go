package service

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	// internal
	model "github.com/sniperkit/colly/pkg/model"
)

// Service represents a service
type Service interface {

	// Login represents...
	Login(ctx context.Context) (string, error)

	// GetEvents represents...
	GetEvents(ctx context.Context, eventChan chan<- *model.EventResult, token, user string, page, count int)

	//-- End
}

// services represents a map of all available interface per service
var services = make(map[string]Service)

// registerService function add a new service in the map/registry of services
func registerService(service Service) {
	services[Name(service)] = service
}

// Name returns the name of a service
func Name(service Service) string {
	parts := strings.Split(reflect.TypeOf(service).String(), ".")
	return strings.ToLower(parts[len(parts)-1])
}

// ForName returns the service for a given name, or an error if it doesn't exist
func ForName(name string) (Service, error) {
	if service, ok := services[strings.ToLower(name)]; ok {
		return service, nil
	}
	return &NotFound{}, fmt.Errorf("service '%s' not found", name)
}
