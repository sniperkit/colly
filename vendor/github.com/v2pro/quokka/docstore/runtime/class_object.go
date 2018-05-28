package runtime

var ObjectClass = NewObject("keys", func(args []interface{}) interface{} {
	obj := args[0].(*DObject)
	keys := make([]interface{}, 0, len(obj.data))
	for key := range obj.data {
		keys = append(keys, key)
	}
	return NewList(keys...)
})
