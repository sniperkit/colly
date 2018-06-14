package colly

type ContentTypeHTML struct {
	resp *Response
	err  error
}

func (c *ContentTypeHTML) Check(resp *Response) error { return ErrNotImplementedYet }
