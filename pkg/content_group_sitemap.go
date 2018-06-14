package colly

type ContentTypeSitemap struct {
	resp *Response
	err  error
}

func (c *ContentTypeSitemap) Check(resp *Response) error { return ErrNotImplementedYet }
