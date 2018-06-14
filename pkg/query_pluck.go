package colly

// Pluck
type Pluck struct {
	Activators  []string
	Deactivator string
}

// Pluckers
type Pluckers []*Pluck
