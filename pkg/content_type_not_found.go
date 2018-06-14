package colly

type ContentTypeNotFound struct {
	resp *Response
	err  error
}

func (c *ContentTypeNotFound) Check(resp *Response) error { return ErrNotImplementedYet }
