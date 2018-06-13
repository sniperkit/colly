package storage

import (
	"context"
)

// Storage is an interface which handles Collector's internal data,
type Storage interface {

	// Init initializes the storage
	Init(ctx context.Context) error

	// WithConfig represents...
	WithConfig(ctx context.Context) error

	// Info represents...
	Info(ctx context.Context) (string, error)

	// Get represents
	Get(ctx context.Context) (string, error)

	// Set represents...
	Set(ctx context.Context) (string, error)

	// Update represents...
	Update(ctx context.Context) (string, error)

	// Delete represents...
	Delete(ctx context.Context) (string, error)

	// Dump
	Dump(ctx context.Context) error

	// Health
	Health(ctx context.Context) (map[string]bool, error)

	// Close
	Close(ctx context.Context) error

	//-- End
}

// storages represents a map of all available interface per service
var stores = make(map[string]Storage)

// registerService function add a new service in the map/registry of services
func registerStorage(store Storage) {
	stores[Name(store)] = store
}

// Name returns the name of a service
func Name(store Storage) string {
	parts := strings.Split(reflect.TypeOf(store).String(), ".")
	return strings.ToLower(parts[len(parts)-1])
}

// ForName returns the service for a given name, or an error if it doesn't exist
func ForName(name string) (Storage, error) {
	if store, ok := stores[strings.ToLower(name)]; ok {
		return storage, nil
	}
	return &NotFound{}, fmt.Errorf("storage '%s' not found", name)
}
