package runtime

import (
	"github.com/json-iterator/go"
	"reflect"
	"unsafe"
	"io"
	"fmt"
)

// unmarshal will only return following types
// float64
// string
// bool
// *DList
// *DObject

var Json = jsoniter.Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       true, // will lose precession
	ObjectFieldMustBeSimpleString: true, // do not unescape object field
}.Froze()

var DebugJson = jsoniter.Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       true, // will lose precession
	ObjectFieldMustBeSimpleString: true, // do not unescape object field
	SortMapKeys:                   true,
}.Froze()

var emptyInterfaceType = reflect.TypeOf((*interface{})(nil)).Elem()
var dobjectType = reflect.TypeOf((*DObject)(nil)).Elem()
var darrayType = reflect.TypeOf((*DList)(nil)).Elem()

func init() {
	Json.RegisterExtension(&jsonExtension{})
	DebugJson.RegisterExtension(&jsonExtension{})
}

type DObject struct {
	Schema  *Schema `json:"-"`
	data    map[string]interface{}
	updated map[string]interface{}
	patched map[string]interface{}
	removed map[string]interface{}
}

func AsBool(val interface{}) bool {
	if val == nil {
		return false
	}
	switch typedVal := val.(type) {
	case bool:
		return typedVal
	case string:
		return typedVal != ""
	case int:
		return typedVal != 0
	case float64:
		return typedVal != 0
	default:
		// empty [] or {} is considered as true
		return true
	}
}

func AsObj(val interface{}) Object {
	return val.(Object)
}

// supress variable not used error
func BlackHole(obj interface{}) {
}

func Call(callee interface{}, args ...interface{}) interface{} {
	return callee.(func([]interface{}) interface{})(args)
}

func NewObject(kv ...interface{}) *DObject {
	data := map[string]interface{}{}
	for i := 0; i < len(kv); i += 2 {
		data[kv[i].(string)] = kv[i+1]
	}
	return &DObject{
		data: data,
	}
}

type DList struct {
	Schema *Schema `json:"-"`
	data   []interface{}
}

func NewList(data ...interface{}) *DList {
	return &DList{data: data}
}

func Get(obj interface{}, path ...interface{}) interface{} {
	for _, elem := range path {
		switch typedElem := elem.(type) {
		case int:
			obj = obj.(*DList).data[typedElem]
		case string:
			obj = obj.(*DObject).data[typedElem]
		}
	}
	return obj
}


type Object interface {
	Set(key interface{}, value interface{})
	Get(key interface{}) interface{}
}

func (obj *DList) Set(keyObj interface{}, value interface{}) {
	panic("list can not be set by property")
}

func (obj *DList) Get(key interface{}) interface{} {
	switch key {
	case "length":
		return len(obj.data)
	case "push":
		return func(args []interface{}) interface{} {
			obj.data = append(obj.data, args[0])
			return len(obj.data)
		}
	default:
		panic(fmt.Sprintf("list does not have property: %v", key))
	}
}

func (obj *DObject) Get(key interface{}) interface{} {
	return obj.data[key.(string)]
}


type jsonExtension struct {
	jsoniter.DummyExtension
}

func (extension *jsonExtension) CreateDecoder(typ reflect.Type) jsoniter.ValDecoder {
	if typ == emptyInterfaceType {
		return &emptyInterfaceDecoder{}
	}
	return nil
}

func (extension *jsonExtension) CreateEncoder(typ reflect.Type) jsoniter.ValEncoder {
	if typ == dobjectType {
		return &objectEncoder{}
	} else if typ == darrayType {
		return &arrayEncoder{}
	}
	return nil
}

type emptyInterfaceDecoder struct {
}

func (decoder *emptyInterfaceDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	nextValueType := iter.WhatIsNext()
	switch nextValueType {
	case jsoniter.ObjectValue:
		obj := &DObject{}
		iter.ReadVal(&obj.data)
		*(*interface{})(ptr) = obj
	case jsoniter.ArrayValue:
		arr := &DList{}
		iter.ReadVal(&arr.data)
		*(*interface{})(ptr) = arr
	case jsoniter.StringValue:
		*(*interface{})(ptr) = iter.ReadString()
	case jsoniter.BoolValue:
		*(*interface{})(ptr) = iter.ReadBool()
	case jsoniter.NilValue:
		iter.ReadNil()
		*(*interface{})(ptr) = nil
	case jsoniter.NumberValue:
		number := iter.SkipAndReturnBytes()
		subIter := Json.BorrowIterator(number)
		intVal := subIter.ReadInt()
		if subIter.Error == io.EOF {
			Json.ReturnIterator(subIter)
			*(*interface{})(ptr) = intVal
		} else {
			Json.ReturnIterator(subIter)
			subIter = Json.BorrowIterator(number)
			*(*interface{})(ptr) = subIter.ReadFloat64()
			if subIter.Error != nil && subIter.Error != io.EOF {
				iter.ReportError("readNumber", subIter.Error.Error())
			}
			Json.ReturnIterator(subIter)
		}
	default:
		iter.ReportError("readEmptyInterface", fmt.Sprintf("unexpected value type: %v", nextValueType))
	}
}

type objectEncoder struct {
}

func (encoder *objectEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	return false
}

func (encoder *objectEncoder) EncodeInterface(val interface{}, stream *jsoniter.Stream) {
	jsoniter.WriteToStream(val, stream, encoder)
}

func (encoder *objectEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DObject)(ptr)
	stream.WriteVal(obj.data)
}

type arrayEncoder struct {
}

func (encoder *arrayEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	return false
}

func (encoder *arrayEncoder) EncodeInterface(val interface{}, stream *jsoniter.Stream) {
	jsoniter.WriteToStream(val, stream, encoder)
}

func (encoder *arrayEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DList)(ptr)
	stream.WriteVal(obj.data)
}

func Dump(obj interface{}) string {
	// DebugJson sorts map key, so it is stable to compare
	encoded, err := DebugJson.MarshalToString(obj)
	if err != nil {
		panic(err)
	}
	return encoded
}
