package json_utils

import "encoding/json"

func Marshal(v interface{}) string {
	jsons, _ := json.Marshal(v)
	return string(jsons)
}

func Unmarshal(data []byte, v interface{}) {
	json.Unmarshal(data, v)
}
