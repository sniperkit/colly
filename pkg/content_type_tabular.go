package colly

// TAB is the key for the tab compatible encodings
const TAB = "tab"

// NewTABDecoder return the right TAB decoder
func NewTABDecoder(isCollection bool) Decoder {
	if isCollection {
		return TABCollectionDecoder // Databook
	}
	return TABDecoder // Dataset
}

// TABDecoder implements the Decoder interface
func TABDecoder(r io.Reader, v *map[string]interface{}) error {
	// *v = mv
	return nil
}

// TABCollectionDecoder implements the Decoder interface over a collection
func TABCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	// *(v) = map[string]interface{}{"collection": mv}
	return nil
}

// ContentTypeTAB...
type ContentTypeTAB struct {
	resp *Response
	err  error
}

// Check...
func (c *ContentTypeTAB) Check(resp *Response) error { return ErrNotImplementedYet }
