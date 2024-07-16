package tq

import (
	"bytes"
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
func unflattenJSONMap(flatMap jsonMap) (nestedMap jsonMap, err error) {
	keys := maps.Keys(flatMap)
	nestedMap = make(jsonMap)
	for len(keys) > 0 {
		key := keys[0]
		if sep := strings.IndexAny(key, ".[]"); sep == -1 {
			nestedMap[key] = flatMap[key]
			keys = slicesRemoveOne(keys, key)
		} else {
			prefix := key[0:sep]
			flatMapPart := make([]jsonMap, 0)
			for key, value := range flatMap {
				if strings.HasPrefix(key, prefix) {
					index := 0
					subkey := strings.TrimPrefix(key, prefix)
					if i, _subkey, found := strings.Cut(subkey, "]"); found && !strings.Contains(i, ".") {
						index, _ = strconv.Atoi(i)
						subkey = _subkey
					}
					flatMapPart[index][subkey] = value
					keys = slicesRemoveOne(keys, key)
				}
			}
			unflat := make([]json.RawMessage, len(flatMapPart))
			for i := range flatMapPart {
				unflat_i, err := unflattenJSONMap(flatMapPart[i])
				if err != nil {
					return nil, err
				}
				unflat[i], err = json.Marshal(unflat_i)
				if err != nil {
					return nil, err
				}
			}
			if len(unflat) > 1 {
				nestedMap[prefix], err = json.Marshal(unflat)
			} else {
				nestedMap[prefix] = unflat[0]
			}
		}
	}
	return
}

func slicesRemove[S any](slice []S, remove []S) (cleanedSlice []S) {
	keep := make(map[any]bool, len(slice)+len(remove))
	for _, item := range slice {
		keep[item] = true
	}
	for _, item := range remove {
		keep[item] = false
	}
	for key, value := range keep {
		if value {
			cleanedSlice = append(cleanedSlice, key.(S))
		}
	}
	return
}

func slicesRemoveOne[S any](slice []S, remove S) (cleanedSlice []S) {
	keep := make(map[any]bool, len(slice)+1)
	for _, item := range slice {
		keep[item] = true
	}
	keep[remove] = false
	for key, value := range keep {
		if value {
			cleanedSlice = append(cleanedSlice, key.(S))
		}
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

func jsonFromMap(in jsonMap) (out []byte, err error) {
	return json.Marshal(in)
}
