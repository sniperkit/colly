package colly

type ContentTypeTAB struct {
	resp *Response
	err  error
}

func (c *ContentTypeTAB) Check(resp *Response) error { return ErrNotImplementedYet }
