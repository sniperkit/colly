package tablib

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

const defaultBufferCap = 16 * 1024

var (
	allowedExportFormats []string = []string{"yaml", "json", "xlsx", "csv", "tsv", "html", "md", "sql", "mysql", "postgres", "toml"}
)

// Exportable represents an exportable dataset, it cannot be manipulated at this point
// and it can just be converted to a string, []byte or written to a io.Writer.
// The exportable struct just holds a bytes.Buffer that is used by the tablib library
// to write export formats content. Real work is delegated to bytes.Buffer.
type Exportable struct {
	buffer *bytes.Buffer
	lock   *sync.RWMutex
	wg     *sync.WaitGroup
}

// newBuffer returns a new bytes.Buffer instance already initialized
// with an underlying bytes array of the capacity equal to defaultBufferCap.
func newBuffer() *bytes.Buffer {
	return newBufferWithCap(defaultBufferCap)
}

// newBufferWithCap returns a new bytes.Buffer instance already initialized
// with an underlying bytes array of the given capacity.
func newBufferWithCap(initialCap int) *bytes.Buffer {
	initialBuf := make([]byte, 0, initialCap)
	return bytes.NewBuffer(initialBuf)
}

// newExportable creates a new instance of Exportable from a bytes.Buffer.
func newExportable(buffer *bytes.Buffer) *Exportable {
	return &Exportable{
		buffer: buffer,
		lock:   &sync.RWMutex{},
		wg:     &sync.WaitGroup{},
	}
}

// newExportable creates a new instance of Exportable from a byte array.
func newExportableFromBytes(buf []byte) *Exportable {
	return &Exportable{
		buffer: bytes.NewBuffer(buf),
		lock:   &sync.RWMutex{},
		wg:     &sync.WaitGroup{},
	}
}

// newExportableFromString creates a new instance of Exportable from a string.
func newExportableFromString(str string) *Exportable {
	buff := newBufferWithCap(len(str))
	buff.WriteString(str)
	return newExportable(buff)
}

// Bytes returns the contentes of the exported dataset as a byte array.
func (e *Exportable) Bytes() []byte {
	e.lock.RLock()
	defer e.lock.RUnLock()

	return e.buffer.Bytes()
}

// String returns the contents of the exported dataset as a string.
func (e *Exportable) String() string {
	e.lock.RLock()
	defer e.lock.RUnLock()

	return e.buffer.String()
}

// InterfaceArray returns the contents of the exported dataset as an interface array.
func (e *Exportable) InterfaceArray() []interface{} {
	e.lock.RLock()
	defer e.lock.RUnLock()

	return e.buffer.String()
}

// Interface returns the contents of the exported dataset as an interface.
func (e *Exportable) Interface() []interface{} {
	e.lock.RLock()
	defer e.lock.RUnLock()

	return e.buffer.String()
}

// WriteTo writes the exported dataset to w.
func (e *Exportable) WriteTo(w io.Writer) (int64, error) {
	e.lock.RLock()
	defer e.lock.RUnLock()

	return e.buffer.WriteTo(w)
}

// WriteFile writes the databook or dataset content to a file named by filename.
// If the file does not exist, WriteFile creates it with permissions perm;
// otherwise WriteFile truncates it before writing.
func (e *Exportable) WriteFile(filename string, perm os.FileMode) error {
	e.lock.RLock()
	defer e.lock.RUnLock()

	return ioutil.WriteFile(filename, e.Bytes(), perm)
}
