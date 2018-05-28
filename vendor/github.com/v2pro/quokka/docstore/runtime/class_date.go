package runtime

import "time"

var DateClass = NewObject()

func init() {
	DateClass.data["now"] = func(args []interface{}) interface{} {
		mockedNow := DateClass.data["mocked_now"]
		if mockedNow != nil {
			return mockedNow
		}
		return int(time.Duration(time.Now().UnixNano()) / time.Millisecond)
	}
}
