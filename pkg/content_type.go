package colly

// HandleOnType struct...
type HandleOnType interface {
	Check(resp *Response) string
}

// handleOnType function...
func (c *Collector) handleOnType(resp *Response) error { return ErrNotImplementedYet }
