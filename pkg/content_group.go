package colly

type HandleOnGroup interface {
	Detect(resp *Response) string
}

func (c *Collector) handleOnGroup(resp *Response) error { return ErrNotImplementedYet }
