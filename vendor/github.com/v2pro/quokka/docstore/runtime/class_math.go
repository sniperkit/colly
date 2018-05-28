package runtime

import (
	"math"
	"reflect"
	"math/rand"
)

var MathClass = NewObject("floor", func(args []interface{}) interface{} {
	switch typedArg := args[0].(type) {
	case float64:
		return int(math.Floor(typedArg))
	case int:
		return typedArg
	default:
		panic("floor does not support type: " + reflect.TypeOf(typedArg).String())
	}
})

func init() {
	MathClass.data["random"] = func(args []interface{}) interface{} {
		mockedRandom := MathClass.data["mocked_random"]
		if mockedRandom != nil {
			return mockedRandom
		}
		return float64(rand.Int31n(1024 * 1024)) / float64(1024 * 1024)
	}
}
