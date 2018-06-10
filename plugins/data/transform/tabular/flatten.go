package tablib

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	// internal
	jsonpointer "github.com/sniperkit/colly/plugins/data/transform/tabular/jsonpointer"
)

var jsonNumberType = reflect.TypeOf(json.Number(""))

func ToMapSlices(data interface{}) ([]map[string]interface{}, error) {
	results := make([]map[string]interface{}, 0)
	v := valueOf(data)
	//v := reflect.ValueOf(data)

	switch v.Kind() {
	case reflect.Map:
		if v.Len() > 0 {
			result, err := flattenMap(v)
			if err != nil {
				return nil, err
			}
			results = append(results, result)
		}

	case reflect.Slice:
		if isObjectArray(v) {
			for i := 0; i < v.Len(); i++ {
				result, err := flattenMap(v.Index(i))
				if err != nil {
					return nil, err
				}
				results = append(results, result)
			}
		} else if v.Len() > 0 {
			result, err := flattenMap(v)
			if err != nil {
				return nil, err
			}
			if result != nil {
				results = append(results, result)
			}
		}

	default:
		return nil, errors.New("Unsupported JSON structure.")
	}

	return results, nil
}

func flatten(obj interface{}) (KeyValue, error) {
	f := make(KeyValue, 0)
	key := jsonpointer.JSONPointer{}
	if err := _flatten(f, obj, key); err != nil {
		return nil, err
	}
	return f, nil
}

func flattenMap(obj interface{}) (map[string]interface{}, error) {
	f := make(map[string]interface{}, 0)
	key := jsonpointer.JSONPointer{}
	if err := _flatten(f, obj, key); err != nil {
		return nil, err
	}
	return f, nil
}

func _flatten(out KeyValue, obj interface{}, key jsonpointer.JSONPointer) error {
	value, ok := obj.(reflect.Value)
	if !ok {
		value = reflect.ValueOf(obj)
	}

	for value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	if value.IsValid() {
		vt := value.Type()
		if vt.AssignableTo(jsonNumberType) {
			out[key.String()] = value.Interface().(json.Number)
			return nil
		}
	}

	switch value.Kind() {
	case reflect.Map:
		_flattenMap(out, value, key)

	case reflect.Slice:
		_flattenSlice(out, value, key)

	case reflect.Ptr:
		if !value.IsNil() {
			out[key.String()] = value.Elem()
		} else {
			out[key.String()] = ""
		}

	case reflect.String:
		out[key.String()] = value.String()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		out[key.String()] = value.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		out[key.String()] = value.Uint()

	case reflect.Float32, reflect.Float64:
		out[key.String()] = value.Float()

	case reflect.Bool:
		out[key.String()] = value.Bool()

	default:
		return fmt.Errorf("Unknown kind: %s", value.Kind())

	}
	return nil
}

/*
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
*/

func _flattenMap(out map[string]interface{}, value reflect.Value, prefix jsonpointer.JSONPointer) {
	keys := sortedMapKeys(value)
	for _, key := range keys {
		pointer := prefix.Clone()
		pointer.AppendString(key.String())
		_flatten(out, value.MapIndex(key).Interface(), pointer)
	}
}

func _flattenSlice(out map[string]interface{}, value reflect.Value, prefix jsonpointer.JSONPointer) {
	for i := 0; i < value.Len(); i++ {
		pointer := prefix.Clone()
		pointer.AppendString(strconv.Itoa(i))
		_flatten(out, value.Index(i).Interface(), pointer)
	}
}

func isObjectArray(obj interface{}) bool {
	value := valueOf(obj)
	if value.Kind() != reflect.Slice {
		return false
	}

	len := value.Len()
	if len == 0 {
		return false
	}
	for i := 0; i < len; i++ {
		if valueOf(value.Index(i)).Kind() != reflect.Map {
			return false
		}
	}
	return true
}
