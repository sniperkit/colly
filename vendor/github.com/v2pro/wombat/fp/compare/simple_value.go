package compare

import "github.com/v2pro/wombat/generic"

func init() {
	ByItself.ImportFunc(compareSimpleValue)
}

var compareSimpleValue = generic.DefineFunc("CompareSimpleValue(val1 T, val2 T) int").
	Param("T", "the type of value to compare").
	Source(`
if val1 < val2 {
	return -1
} else if val1 == val2 {
	return 0
} else {
	return 1
}`)
