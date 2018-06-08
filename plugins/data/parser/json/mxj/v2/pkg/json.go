// Copyright 2012-2014 Charles Banning. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

package mxj

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// ------------------------------ write JSON -----------------------

// Just a wrapper on json.Marshal.
// If option safeEncoding is'true' then safe encoding of '<', '>' and '&'
// is preserved. (see encoding/json#Marshal, encoding/json#Encode)
func (mv Map) Json(safeEncoding ...bool) ([]byte, error) {
	var s bool
	if len(safeEncoding) == 1 {
		s = safeEncoding[0]
	}

	b, err := json.Marshal(mv)

	if !s {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

// Just a wrapper on json.MarshalIndent.
// If option safeEncoding is'true' then safe encoding of '<' , '>' and '&'
// is preserved. (see encoding/json#Marshal, encoding/json#Encode)
func (mv Map) JsonIndent(prefix, indent string, safeEncoding ...bool) ([]byte, error) {
	var s bool
	if len(safeEncoding) == 1 {
		s = safeEncoding[0]
	}

	b, err := json.MarshalIndent(mv, prefix, indent)
	if !s {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

// The following implementation is provided for symmetry with NewMapJsonReader[Raw]
// The names will also provide a key for the number of return arguments.

// Writes the Map as JSON on the Writer.
// If 'safeEncoding' is 'true', then "safe" encoding of '<', '>' and '&' is preserved.
func (mv Map) JsonWriter(jsonWriter io.Writer, safeEncoding ...bool) error {
	b, err := mv.Json(safeEncoding...)
	if err != nil {
		return err
	}

	_, err = jsonWriter.Write(b)
	return err
}

// Writes the Map as JSON on the Writer. []byte is the raw JSON that was written.
// If 'safeEncoding' is 'true', then "safe" encoding of '<', '>' and '&' is preserved.
func (mv Map) JsonWriterRaw(jsonWriter io.Writer, safeEncoding ...bool) ([]byte, error) {
	b, err := mv.Json(safeEncoding...)
	if err != nil {
		return b, err
	}

	_, err = jsonWriter.Write(b)
	return b, err
}

// Writes the Map as pretty JSON on the Writer.
// If 'safeEncoding' is 'true', then "safe" encoding of '<', '>' and '&' is preserved.
func (mv Map) JsonIndentWriter(jsonWriter io.Writer, prefix, indent string, safeEncoding ...bool) error {
	b, err := mv.JsonIndent(prefix, indent, safeEncoding...)
	if err != nil {
		return err
	}

	_, err = jsonWriter.Write(b)
	return err
}

// Writes the Map as pretty JSON on the Writer. []byte is the raw JSON that was written.
// If 'safeEncoding' is 'true', then "safe" encoding of '<', '>' and '&' is preserved.
func (mv Map) JsonIndentWriterRaw(jsonWriter io.Writer, prefix, indent string, safeEncoding ...bool) ([]byte, error) {
	b, err := mv.JsonIndent(prefix, indent, safeEncoding...)
	if err != nil {
		return b, err
	}

	_, err = jsonWriter.Write(b)
	return b, err
}

// --------------------------- read JSON -----------------------------

// Decode numericvalues as json.Number type Map values - see encoding/json#Number.
// NOTE: this is for decoding JSON into a Map with NewMapJson(), NewMapJsonReader(),
// etc.; it does not affect NewMapXml(), etc.  The XML encoders mv.Xml() and mv.XmlIndent()
// do recognize json.Number types; a JSON object can be decoded to a Map with json.Number
// value types and the resulting Map can be correctly encoded into a XML object.
var JsonUseNumber bool = false

// Just a wrapper on json.Unmarshal
//	Converting JSON to XML is a simple as:
//		...
//		mapVal, merr := mxj.NewMapJson(jsonVal)
//		if merr != nil {
//			// handle error
//		}
//		xmlVal, xerr := mapVal.Xml()
//		if xerr != nil {
//			// handle error
//		}
// NOTE: as a special case, passing a list, e.g., [{"some-null-value":"", "a-non-null-value":"bar"}],
// will be interpreted as having the root key 'object' prepended - {"object":[ ... ]} - to unmarshal to a Map.
// See mxj/j2x/j2x_test.go.
func NewMapJson(jsonVal []byte) (Map, error) {
	// empty or nil begets empty
	if len(jsonVal) == 0 {
		m := make(map[string]interface{}, 0)
		return m, nil
	}
	// handle a goofy case ...
	if jsonVal[0] == '[' {
		jsonVal = []byte(`{"array":` + string(jsonVal) + `}`)
	}
	m := make(map[string]interface{})
	buf := bytes.NewReader(jsonVal)
	dec := json.NewDecoder(buf)
	if JsonUseNumber {
		dec.UseNumber()
	}
	err := dec.Decode(&m)
	return m, err
}

func NewMapJsonArray(jsonVal []byte) (Map, error) {
	var items []interface{}
	// empty or nil begets empty
	if len(jsonVal) == 0 {
		empty := make(map[string]interface{}, 0)
		return empty, nil
	}

	buf := bytes.NewReader(jsonVal)
	dec := json.NewDecoder(buf)
	if JsonUseNumber {
		dec.UseNumber()
	}
	if err := dec.Decode(&items); err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	return map[string]interface{}{"entries": items}, nil
}

/*
// NewJSONDecoder return the right JSON decoder
func NewJSONDecoder(isCollection bool) Decoder {
	if isCollection {
		return JSONCollectionDecoder
	}
	return JSONDecoder
}

// JSONDecoder implements the Decoder interface
func JSONDecoder(r io.Reader, v *map[string]interface{}) error {
	d := json.NewDecoder(r)
	d.UseNumber()
	return d.Decode(v)
}

// JSONCollectionDecoder implements the Decoder interface over a collection
func JSONCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	var collection []interface{}
	d := json.NewDecoder(r)
	d.UseNumber()
	if err := d.Decode(&collection); err != nil {
		return err
	}
	*(v) = map[string]interface{}{"collection": collection}
	return nil
}
*/

// Retrieve a Map value from an io.Reader.
//  NOTE: The raw JSON off the reader is buffered to []byte using a ByteReader. If the io.Reader is an
//        os.File, there may be significant performance impact. If the io.Reader is wrapping a []byte
//        value in-memory, however, such as http.Request.Body you CAN use it to efficiently unmarshal
//        a JSON object.
func NewMapJsonReaderAll(jsonReader io.Reader) (Map, error) {
	dec := json.NewDecoder(jsonReader)
	if JsonUseNumber {
		dec.UseNumber()
	}
	m := make(map[string]interface{})
	err := dec.Decode(&m) // Unmarshal the 'presumed' JSON string
	if err != nil {
		return nil, err
	}
	return m, nil
}

func NewMapJsonArrayReaderAll(jsonReader io.Reader) (Map, error) {
	var items []interface{}
	dec := json.NewDecoder(jsonReader)
	if JsonUseNumber {
		dec.UseNumber()
	}
	if err := dec.Decode(&items); err != nil {
		return nil, err
	}
	// m := map[string]interface{}{"items": items}
	return map[string]interface{}{"items": items}, nil
}

// Retrieve a Map value from an io.Reader.
//  NOTE: The raw JSON off the reader is buffered to []byte using a ByteReader. If the io.Reader is an
//        os.File, there may be significant performance impact. If the io.Reader is wrapping a []byte
//        value in-memory, however, such as http.Request.Body you CAN use it to efficiently unmarshal
//        a JSON object.
func NewMapJsonReader(jsonReader io.Reader) (Map, error) {
	jb, err := getJson(jsonReader)
	if err != nil || len(*jb) == 0 {
		return nil, err
	}

	// Unmarshal the 'presumed' JSON string
	return NewMapJson(*jb)
}

// Retrieve a Map value and raw JSON - []byte - from an io.Reader.
//  NOTE: The raw JSON off the reader is buffered to []byte using a ByteReader. If the io.Reader is an
//        os.File, there may be significant performance impact. If the io.Reader is wrapping a []byte
//        value in-memory, however, such as http.Request.Body you CAN use it to efficiently unmarshal
//        a JSON object and retrieve the raw JSON in a single call.
func NewMapJsonReaderRaw(jsonReader io.Reader) (Map, []byte, error) {
	jb, err := getJson(jsonReader)
	if err != nil || len(*jb) == 0 {
		return nil, *jb, err
	}

	// Unmarshal the 'presumed' JSON string
	m, merr := NewMapJson(*jb)
	return m, *jb, merr
}

// Pull the next JSON string off the stream: just read from first '{' to its closing '}'.
// Returning a pointer to the slice saves 16 bytes - maybe unnecessary, but internal to package.
func getJson(rdr io.Reader) (*[]byte, error) {
	bval := make([]byte, 1)
	jb := make([]byte, 0)
	var inQuote, inJson bool
	var parenCnt int
	var previous byte

	// scan the input for a matched set of {...}
	// json.Unmarshal will handle syntax checking.
	for {
		_, err := rdr.Read(bval)
		if err != nil {
			if err == io.EOF && inJson && parenCnt > 0 {
				return &jb, fmt.Errorf("no closing } for JSON string: %s", string(jb))
			}
			return &jb, err
		}
		switch bval[0] {
		case '{':
			if !inQuote {
				parenCnt++
				inJson = true
			}
		case '}':
			if !inQuote {
				parenCnt--
			}
			if parenCnt < 0 {
				return nil, fmt.Errorf("closing } without opening {: %s", string(jb))
			}
		case '"':
			if inQuote {
				if previous == '\\' {
					break
				}
				inQuote = false
			} else {
				inQuote = true
			}
		case '\n', '\r', '\t', ' ':
			if !inQuote {
				continue
			}
		}
		if inJson {
			jb = append(jb, bval[0])
			if parenCnt == 0 {
				break
			}
		}
		previous = bval[0]
	}

	return &jb, nil
}

// ------------------------------- JSON Reader handler via Map values  -----------------------

// Default poll delay to keep Handler from spinning on an open stream
// like sitting on os.Stdin waiting for imput.
var jhandlerPollInterval = time.Duration(1e6)

// While unnecessary, we make HandleJsonReader() have the same signature as HandleXmlReader().
// This avoids treating one or other as a special case and discussing the underlying stdlib logic.

// Bulk process JSON using handlers that process a Map value.
//	'rdr' is an io.Reader for the JSON (stream).
//	'mapHandler' is the Map processing handler. Return of 'false' stops io.Reader processing.
//	'errHandler' is the error processor. Return of 'false' stops io.Reader  processing and returns the error.
//	Note: mapHandler() and errHandler() calls are blocking, so reading and processing of messages is serialized.
//	      This means that you can stop reading the file on error or after processing a particular message.
//	      To have reading and handling run concurrently, pass argument to a go routine in handler and return 'true'.
func HandleJsonReader(jsonReader io.Reader, mapHandler func(Map) bool, errHandler func(error) bool) error {
	var n int
	for {
		m, merr := NewMapJsonReader(jsonReader)
		n++

		// handle error condition with errhandler
		if merr != nil && merr != io.EOF {
			merr = fmt.Errorf("[jsonReader: %d] %s", n, merr.Error())
			if ok := errHandler(merr); !ok {
				// caused reader termination
				return merr
			}
			continue
		}

		// pass to maphandler
		if len(m) != 0 {
			if ok := mapHandler(m); !ok {
				break
			}
		} else if merr != io.EOF {
			<-time.After(jhandlerPollInterval)
		}

		if merr == io.EOF {
			break
		}
	}
	return nil
}

// Bulk process JSON using handlers that process a Map value and the raw JSON.
//	'rdr' is an io.Reader for the JSON (stream).
//	'mapHandler' is the Map and raw JSON - []byte - processor. Return of 'false' stops io.Reader processing.
//	'errHandler' is the error and raw JSON processor. Return of 'false' stops io.Reader processing and returns the error.
//	Note: mapHandler() and errHandler() calls are blocking, so reading and processing of messages is serialized.
//	      This means that you can stop reading the file on error or after processing a particular message.
//	      To have reading and handling run concurrently, pass argument(s) to a go routine in handler and return 'true'.
func HandleJsonReaderRaw(jsonReader io.Reader, mapHandler func(Map, []byte) bool, errHandler func(error, []byte) bool) error {
	var n int
	for {
		m, raw, merr := NewMapJsonReaderRaw(jsonReader)
		n++

		// handle error condition with errhandler
		if merr != nil && merr != io.EOF {
			merr = fmt.Errorf("[jsonReader: %d] %s", n, merr.Error())
			if ok := errHandler(merr, raw); !ok {
				// caused reader termination
				return merr
			}
			continue
		}

		// pass to maphandler
		if len(m) != 0 {
			if ok := mapHandler(m, raw); !ok {
				break
			}
		} else if merr != io.EOF {
			<-time.After(jhandlerPollInterval)
		}

		if merr == io.EOF {
			break
		}
	}
	return nil
}
