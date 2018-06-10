package tablib

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	reflectx "github.com/corpix/reflect"
)

// getBytes
func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// inArray
func inArray(input string, array []string) bool {
	for _, v := range array {
		if input == v {
			return true
		}
	}
	return false
}

// internalLoadFromDict creates a Dataset from an array of map representing columns.
func internalLoadFromDict(input []map[string]interface{}) (*Dataset, error) {
	// retrieve columns
	headers := make([]string, 0, 10)
	for h := range input[0] {
		headers = append(headers, h)
	}

	ds := NewDataset(headers)
	for _, e := range input {
		row := make([]interface{}, 0, len(headers))
		for _, h := range headers {
			row = append(row, e[h])
		}
		ds.AppendValues(row...)
	}

	return ds, nil
}

// isTagged checks if a tag is in an array of tags.
func isTagged(tag string, tags []string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

// asString returns a value as a string.
func (d *Dataset) asString(vv interface{}) string {
	var v string
	switch vv.(type) {
	case string:
		v = vv.(string)
	case int:
		v = strconv.Itoa(vv.(int))
	case int64:
		v = strconv.FormatInt(vv.(int64), 10)
	case uint64:
		v = strconv.FormatUint(vv.(uint64), 10)
	case bool:
		v = strconv.FormatBool(vv.(bool))
	case float64:
		v = strconv.FormatFloat(vv.(float64), 'G', -1, 32)
	case json.Number:
		v = vv.(json.Number).String()
	case time.Time:
		v = vv.(time.Time).Format(time.RFC3339)
	default:
		if d.EmptyValue != "" {
			v = d.EmptyValue
		} else {
			v = fmt.Sprintf("%s", v)
		}
	}
	return v
}

func valueOf(obj interface{}) reflect.Value {
	v, ok := obj.(reflect.Value)
	if !ok {
		v = reflect.ValueOf(obj)
	}

	for v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
		if v.IsValid() {
			v = v.Elem()
		} else {
			v.String()
		}
	}

	return v
}

func toDebug(obj interface{}) string {

	return ""

	v, ok := obj.(reflect.Value)
	if !ok {
		v = reflect.ValueOf(obj)
	}

	var debugMsg string

	indirectValue := reflectx.IndirectValue(v)
	indirectType := reflectx.IndirectType(reflect.TypeOf(v))

	switch v.Kind() {
	case reflect.String:
		debugMsg = fmt.Sprintln("reflect.String: Kind=", v.Kind(), ",TypeOf=", typeOf(obj), "IsValid=", v.IsValid())
	case reflect.Map:
		debugMsg = fmt.Sprintln("reflect.Map: Kind=", v.Kind(), ",TypeOf=", typeOf(obj), "indirectType=", indirectType, "IsNil=", v.IsNil(), "IsValid=", v.IsValid(), "len=", v.Len(), ",Elem=", v.Elem())
	case reflect.Ptr:
		debugMsg = fmt.Sprintln("reflect.Ptr: Kind=", v.Kind(), "indirectValue=", indirectValue, "indirectType=", indirectType, ",TypeOf=", typeOf(obj), "IsNil=", v.IsNil(), "IsValid=", v.IsValid(), ",Elem=", v.Elem())
	case reflect.Interface:
		debugMsg = fmt.Sprintln("reflect.Interface: Kind=", v.Kind(), ",TypeOf=", typeOf(obj), "indirectValue=", indirectValue, "indirectType=", indirectType, "IsNil=", v.IsNil(), "IsValid=", v.IsValid(), ",Elem=", v.Elem())

	default:
		debugMsg = fmt.Sprintln("reflect.default: Kind=", v.Kind(), "IsSimpleType=", IsSimpleType(v.Kind()), ",TypeOf=", typeOf(obj), "IsValid=", v.IsValid())
	}

	return strings.Replace(debugMsg, "\n", "", -1)
}

func toString(obj interface{}) string {
	return fmt.Sprintf("%v", obj)
	//return fmt.Sprintf("%v|%s", obj, toDebug(obj))
}

func fmtTypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func typeOf(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func IsSimpleType(kind reflect.Kind) bool {
	switch kind {
	case reflect.Bool:
		return true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Float32, reflect.Float64:
		return true
	case reflect.String:
		return true
	}
	return false
}

func sortedKeys(m map[string]interface{}) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func InSlice(needle interface{}, haystack []interface{}) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// from http://stackoverflow.com/questions/19374219/how-to-find-the-difference-between-two-slices-of-strings-in-golang
func SliceDiff(slice1 []string, slice2 []string) []string {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
	return diff
}
