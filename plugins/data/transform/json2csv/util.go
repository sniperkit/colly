package json2csv

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	reflectx "github.com/corpix/reflect"
)

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
