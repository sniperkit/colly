package colly

// HandleOnType struct...
type HandleOnType interface {
	Check(resp *Response) string
}

// handleOnType function...
func (c *Collector) handleOnType(resp *Response) error { return ErrNotImplementedYet }

/*
func (s *Collector) initEncodingDefaults() {
	switch strings.ToLower(collector.Encoding) {
	case TAB:
		collector.Decoder = NewTABDecoder(false)
	case HTML:
		collector.Decoder = NewHTMLDecoder(false)
	case XML:
		collector.Decoder = NewXMLDecoder(false)
	case RSS:
		collector.Decoder = NewRSSDecoder(false)
	default:
		collector.Decoder = NewJSONDecoder(false)
	}
}
*/
