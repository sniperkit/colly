package colly

// Extractors
type Extractors []*Extractor

// Extractor
type Extractor struct {
	Activators  []string
	Deactivator string
}

// Pluckers
type Pluckers []*Pluck

// Pluck
type Pluck struct {
	Activators  []string
	Deactivator string
}
