package colly

type ContentTypeJSON struct {
	resp *Response
	err  error
}

func (c *ContentTypeJSON) Check(resp *Response) error { return ErrNotImplementedYet }
