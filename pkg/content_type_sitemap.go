package colly

// ContentTypeSitemap...
type ContentTypeSitemap struct {
	resp *Response
	err  error
}

// Check...
func (c *ContentTypeSitemap) Check(resp *Response) error { return ErrNotImplementedYet }
