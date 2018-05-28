package runtime

import "reflect"

func GT(left interface{}, right interface{}) bool {
	switch typedLeft := left.(type) {
	case int:
		switch typedRight := right.(type) {
		case int:
			return int(typedLeft) > typedRight
		case float64:
			return float64(typedLeft) > typedRight
		default:
			panic("can not compare int with " + reflect.TypeOf(right).String())
		}
	case float64:
		switch typedRight := right.(type) {
		case int:
			return typedLeft > float64(typedRight)
		case float64:
			return typedLeft > typedRight
		default:
			panic("can not compare float with " + reflect.TypeOf(right).String())
		}
	default:
		panic("compare does not support " + reflect.TypeOf(left).String())
	}
}

func GE(left interface{}, right interface{}) bool {
	switch typedLeft := left.(type) {
	case int:
		switch typedRight := right.(type) {
		case int:
			return int(typedLeft) >= typedRight
		case float64:
			return float64(typedLeft) >= typedRight
		default:
			panic("can not compare int with " + reflect.TypeOf(right).String())
		}
	case float64:
		switch typedRight := right.(type) {
		case int:
			return typedLeft >= float64(typedRight)
		case float64:
			return typedLeft >= typedRight
		default:
			panic("can not compare float with " + reflect.TypeOf(right).String())
		}
	default:
		panic("compare does not support " + reflect.TypeOf(left).String())
	}
}

func LT(left interface{}, right interface{}) bool {
	switch typedLeft := left.(type) {
	case int:
		switch typedRight := right.(type) {
		case int:
			return int(typedLeft) < typedRight
		case float64:
			return float64(typedLeft) < typedRight
		default:
			panic("can not compare int with " + reflect.TypeOf(right).String())
		}
	case float64:
		switch typedRight := right.(type) {
		case int:
			return typedLeft < float64(typedRight)
		case float64:
			return typedLeft < typedRight
		default:
			panic("can not compare float with " + reflect.TypeOf(right).String())
		}
	default:
		panic("compare does not support " + reflect.TypeOf(left).String())
	}
}

func LE(left interface{}, right interface{}) bool {
	switch typedLeft := left.(type) {
	case int:
		switch typedRight := right.(type) {
		case int:
			return int(typedLeft) <= typedRight
		case float64:
			return float64(typedLeft) <= typedRight
		default:
			panic("can not compare int with " + reflect.TypeOf(right).String())
		}
	case float64:
		switch typedRight := right.(type) {
		case int:
			return typedLeft <= float64(typedRight)
		case float64:
			return typedLeft <= typedRight
		default:
			panic("can not compare float with " + reflect.TypeOf(right).String())
		}
	default:
		panic("compare does not support " + reflect.TypeOf(left).String())
	}
}

func EQ(left interface{}, right interface{}) bool {

	switch typedLeft := left.(type) {
	case int:
		switch typedRight := right.(type) {
		case int:
			return int(typedLeft) == typedRight
		case float64:
			return float64(typedLeft) == typedRight
		default:
			panic("can not compare int with " + reflect.TypeOf(right).String())
		}
	case float64:
		switch typedRight := right.(type) {
		case int:
			return typedLeft == float64(typedRight)
		case float64:
			return typedLeft == typedRight
		default:
			panic("can not compare float with " + reflect.TypeOf(right).String())
		}
	case string:
		switch typedRight := right.(type) {
		case string:
			return typedLeft == typedRight
		default:
			panic("can not compare string with " + reflect.TypeOf(right).String())
		}
	default:
		panic("compare does not support " + reflect.TypeOf(left).String())
	}
}