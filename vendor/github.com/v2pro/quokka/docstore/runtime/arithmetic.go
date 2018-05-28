package runtime

import "reflect"

func Add(left interface{}, right interface{}) interface{} {
	switch typedLeft := left.(type) {
	case int:
		switch typedRight := right.(type) {
		case int:
			return int(typedLeft) + typedRight
		case float64:
			return float64(typedLeft) + typedRight
		default:
			panic("can not add int with " + reflect.TypeOf(right).String())
		}
	case float64:
		switch typedRight := right.(type) {
		case int:
			return typedLeft + float64(typedRight)
		case float64:
			return typedLeft + typedRight
		default:
			panic("can not add float with " + reflect.TypeOf(right).String())
		}
	default:
		panic("add does not support " + reflect.TypeOf(left).String())
	}
}

func Subtract(left interface{}, right interface{}) interface{} {
	switch typedLeft := left.(type) {
	case int:
		switch typedRight := right.(type) {
		case int:
			return int(typedLeft) - typedRight
		case float64:
			return float64(typedLeft) - typedRight
		default:
			panic("can not subtract int with " + reflect.TypeOf(right).String())
		}
	case float64:
		switch typedRight := right.(type) {
		case int:
			return typedLeft - float64(typedRight)
		case float64:
			return typedLeft - typedRight
		default:
			panic("can not subtract float with " + reflect.TypeOf(right).String())
		}
	default:
		panic("subtract does not support " + reflect.TypeOf(left).String())
	}
}

func Multiply(left interface{}, right interface{}) interface{} {
	switch typedLeft := left.(type) {
	case int:
		switch typedRight := right.(type) {
		case int:
			return int(typedLeft) * typedRight
		case float64:
			return float64(typedLeft) * typedRight
		default:
			panic("can not multiply int with " + reflect.TypeOf(right).String())
		}
	case float64:
		switch typedRight := right.(type) {
		case int:
			return typedLeft * float64(typedRight)
		case float64:
			return typedLeft * typedRight
		default:
			panic("can not multiply float with " + reflect.TypeOf(right).String())
		}
	default:
		panic("multiply does not support " + reflect.TypeOf(left).String())
	}
}

func Divide(left interface{}, right interface{}) interface{} {
	switch typedLeft := left.(type) {
	case int:
		switch typedRight := right.(type) {
		case int:
			return float64(typedLeft) / float64(typedRight)
		case float64:
			return float64(typedLeft) / typedRight
		default:
			panic("can not divide int with " + reflect.TypeOf(right).String())
		}
	case float64:
		switch typedRight := right.(type) {
		case int:
			return typedLeft / float64(typedRight)
		case float64:
			return typedLeft / typedRight
		default:
			panic("can not divide float with " + reflect.TypeOf(right).String())
		}
	default:
		panic("divide does not support " + reflect.TypeOf(left).String())
	}
}

func NegativeOf(val interface{}) interface{} {
	switch typedVal := val.(type) {
	case int:
		return -typedVal
	case float64:
		return -typedVal
	default:
		panic("negativeOf does not support " + reflect.TypeOf(val).String())
	}
}
