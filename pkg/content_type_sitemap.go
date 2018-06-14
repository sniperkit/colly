package colly

// SITEMAP is the key for the sitemap encoding
const SITEMAP = "sitemap"

// NewSITEMAPDecoder return the right SITEMAP decoder
func NewSITEMAPDecoder(isCollection bool) Decoder {
	if isCollection {
		return SITEMAPCollectionDecoder // Databook
	}
	return SITEMAPDecoder // Dataset
}

// SITEMAPDecoder implements the Decoder interface
func SITEMAPDecoder(r io.Reader, v *map[string]interface{}) error {
	// *v = mv
	return nil
}

// SITEMAPCollectionDecoder implements the Decoder interface over a collection/sitemap index
func SITEMAPCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	// *(v) = map[string]interface{}{"collection": mv}
	return nil
}

// ContentTypeSitemap...
type ContentTypeSitemap struct {
	resp *Response
	err  error
}

// Check...
func (c *ContentTypeSitemap) Check(resp *Response) error { return ErrNotImplementedYet }
