package colly

type HandleOnContent interface {
	Detect(resp *Response) string
}

func (c *Collector) handleOnContent(resp *Response) error { return ErrNotImplementedYet }
