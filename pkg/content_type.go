package colly

type HandleOnType interface {
	Check(resp *Response) string
}

func (c *Collector) handleOnType(resp *Response) error { return ErrNotImplementedYet }
