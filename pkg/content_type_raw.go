package colly

type ContentTypeRAW struct {
	resp *Response
	err  error
}

func (c *ContentTypeRAW) Check(resp *Response) error { return ErrNotImplementedYet }
