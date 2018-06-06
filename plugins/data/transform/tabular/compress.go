package tablib

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

var (
	allowedCompressionFormats []string = []string{"gz", "tar.gz", "rar", "zip", "7z"}
)

// Exportable represents an exportable dataset, it cannot be manipulated at this point
// and it can just be converted to a string, []byte or written to a io.Writer.
// The exportable struct just holds a bytes.Buffer that is used by the tablib library
// to write export formats content. Real work is delegated to bytes.Buffer.
type Compressable struct {
	buffer *bytes.Buffer
	level  int
	lock   *sync.RWMutex
	wg     *sync.WaitGroup
}
