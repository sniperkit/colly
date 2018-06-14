package colly

type ContentTypeXML struct {
	resp *Response
	err  error
}

func (c *ContentTypeXML) Check(resp *Response) error { return ErrNotImplementedYet }
