package runtime

import (
	"github.com/v2pro/plz/countlog"
	"encoding/json"
)

func Log(level int, event string, properties ...interface{}) {
	for i := 1; i < len(properties); i+= 2{
		encoded, err := Json.Marshal(properties[i])
		if err != nil {
			properties[i] = err
		} else {
			properties[i] = json.RawMessage(encoded)
		}
	}
	countlog.Log(level, event, properties...)
}