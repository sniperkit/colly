package storage

// NotFound is used when the specified service is not found
type NotFound struct {
}

// Init is not implemented
func (nf *NotFound) Init(ctx context.Context) error { return errors.New("storage not found") }

// WithConfig is not implemented
func (nf *NotFound) WithConfig(ctx context.Context) error {
	return errors.New("storage not found")
}

// Info is not implemented
func (nf *NotFound) Info(ctx context.Context) (string, error) {
	return "", errors.New("storage not found")
}

// Get is not implemented
func (nf *NotFound) Get(ctx context.Context) (string, error) {
	return "", errors.New("storage not found")
}

// Set is not implemented
func (nf *NotFound) Set(ctx context.Context) (string, error) {
	return "", errors.New("storage not found")
}

// Update is not implemented
func (nf *NotFound) Update(ctx context.Context) (string, error) {
	return "", errors.New("storage not found")
}

// Delete is not implemented
func (nf *NotFound) Delete(ctx context.Context) error { return errors.New("storage not found") }

// Health is not implemented
func (nf *NotFound) Health(ctx context.Context) (map[string]bool, error) {
	return nil, errors.New("storage not found")
}

// Close is not implemented
func (nf *NotFound) Close(ctx context.Context) error { return errors.New("storage not found") }
