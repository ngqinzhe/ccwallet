package util

import "encoding/json"

func SafeJsonDump(x interface{}) string {
	b, err := json.Marshal(x)
	if err != nil {
		return ""
	}
	return string(b)
}
