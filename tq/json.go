package tq

import (
	"encoding/json"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

// String map of json raw strings
type jsonMap map[string]json.RawMessage

// Flattens a jsonMap with possibly nested values to an unnested
// (flat) jsonMap such that each key is the JSONPath location of the value
// in the input object.
//
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
		value = bytes.TrimSpace(value)
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

// Unflattens a jsonMap by parsing each key as the JSONPath location of the
// value in the final jsonMap.
func unflattenJSONMap(flatMap jsonMap) (nestedMap jsonMap) {
	keys := maps.Keys(flatMap)
	nestedMap = make(jsonMap)
	for key, value := range flatMap {
		nestedMapPart := make(jsonMap)
		sep := strings.Index(key, ".")
		if sep == -1 {
			continue
		}
		prefix := key[0:]
	}
	return nestedMap
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

func jsonFromMap(in jsonMap) (out []byte, err error) {
	return json.Marshal(in)
}
