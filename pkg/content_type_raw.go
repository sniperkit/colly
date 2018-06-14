package colly

// ContentTypeRAW...
type ContentTypeRAW struct {
	resp *Response
	err  error
}

// Check...
func (c *ContentTypeRAW) Check(resp *Response) error { return ErrNotImplementedYet }
