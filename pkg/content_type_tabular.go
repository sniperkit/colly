package colly

// ContentTypeTAB...
type ContentTypeTAB struct {
	resp *Response
	err  error
}

// Check...
func (c *ContentTypeTAB) Check(resp *Response) error { return ErrNotImplementedYet }
