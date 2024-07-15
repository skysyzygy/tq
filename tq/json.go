package tq

import (
	"encoding/json"
	"strconv"
)

type jsonMap map[string]json.RawMessage

// Flattens possibly nested JSON object to an unnested object with each
// key defined by the JSONPath location in the input object.
// Assumption:
// All json is built from key/value `<objects>` like:
// `{"key1":<value1>,"key2":<value2>}`
// where `<value>` can be atomic (string/numeric/boolean/null), or
// `<objects>` or arrays of `<objects>`.
func flattenJSONMap(nestedMap jsonMap, prefix string) (flatMap jsonMap,
	err error) {
	flatMap = make(jsonMap)
	if prefix != "" {
		prefix = prefix + "."
	}
	for key, value := range nestedMap {
		flatMapPart := make(jsonMap)
		switch string(value[0]) {
		case "{":
			nestedMapPart, _err := jsonToMap(value)
			if _err != nil {
				return nil, _err
			}
			flatMapPart, err = flattenJSONMap(nestedMapPart, prefix+key)
		case "[":
			array := make([]json.RawMessage, 0, 1024)
			err = json.Unmarshal(value, &array)
			if err != nil {
				return nil, err
			}
			for i, aValue := range array {
				nestedMapPart, err := jsonToMap(aValue)
				if err != nil {
					return nil, err
				}
				flatMapPartPart, err := flattenJSONMap(nestedMapPart, prefix+key+"["+strconv.Itoa(i)+"]")
				if err != nil {
					return nil, err
				}
				updateJSONMap(flatMapPartPart, flatMapPart)
			}
		default:
			flatMapPart = make(jsonMap)
			flatMapPart[prefix+key] = value
		}
		updateJSONMap(flatMapPart, flatMap)
	}
	return
}

func updateJSONMap(in jsonMap,
	out jsonMap) {
	for key, value := range in {
		out[key] = value
	}
}

func jsonToMap(in []byte) (out jsonMap, err error) {
	err = json.Unmarshal(in, &out)
	if err != nil {
		return nil, err
	}
	return
}
