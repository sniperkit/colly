package tablib

type Selector struct {
	PatternURL string
	Headers    []string
	Mixed      []interface{}
	Slicer     *Slicer
}

type Slicer struct {
	Cols string
	Rows string
}
