package colly

type ContentTypeRSS struct {
	resp *Response
	err  error
}

func (c *ContentTypeRSS) Check(resp *Response) error { return ErrNotImplementedYet }
