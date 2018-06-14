package colly

// Query
type Query struct {
	Activators  []string
	Deactivator string
}

// Queries
type Queries []*Query
