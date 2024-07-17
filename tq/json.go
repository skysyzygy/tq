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

func flattenJSONError(key string, err error) error {
		errors.Wrap(errors.New("couldn't flatten))
}

// Flattens a jsonMap with possibly nested values to an unnested
// (flat) jsonMap such that each key is the JSONPath location of the value
// in the input object.
//
// Assumptions:
// There are no duplicate keys 
// No keys contain the characters .,[,]
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
// make json a map
			nestedMapPart := make(jsonMap)
			err = json.Unmarshal(value, &nestedMapPart)
// recurse
			flatMapPart, err = flattenJSONMap(nestedMapPart, prefix+key)
		case "[":
// split json array
			array := make([]json.RawMessage)
			err = json.Unmarshal(value, &array)
			if err != nil {
				return nil, err
			}
			for i, value := range array {
// recurse back to flattenJSON to allow atomic elements
				flatMapPartPart, err := flattenJSON(value, prefix+key+"["+strconv.Itoa(i)+"]")
				if err != nil {
					return nil, err
				}
				updateJSONMap(flatMapPartPart, flatMapPart)
			}
		default:
// handle atomic values
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
// len(keys) is an upper bound
			flatMapPart := make([]jsonMap, 0, len(keys))
			matchedKeys := make([]string, 0, len(keys))
	for len(keys) > 0 {
		key := keys[0]
// handle bare key
		if sep := strings.IndexAny(key, ".[]"); sep == -1 {
			nestedMap[key] = flatMap[key]
			keys = slicesRemove(keys, []string{key})
		} else {
			prefix := key[0:sep]
			for _, key := range keys {
				if strings.HasPrefix(key, prefix) {
value := flatMap[key]
					matchedKeys = append(matchedKeys, key)
					index := 0
					key = strings.TrimPrefix(key, prefix)
					if i, _key, found := strings.Cut(key, "]"); found && !strings.Contains(i, ".") {
						index, err = strconv.Atoi(i)
						key = _key
						if err != nil {
							return nil, err
						}
					}
					flatMapPart[index][key] = value
				}
			}
			keys = slicesRemove(keys, matchedKeys)
			unflat := make([]json.RawMessage, len(flatMapPart))
			for i := range flatMapPart {
// recurse
				_unflat, err := unflattenJSONMap(flatMapPart[i])
				if err != nil {
					return nil, err
				}
// combine into json message
				unflat[i], err = json.Marshal(_unflat)
				if err != nil {
					return nil, err
				}
			}
			if len(unflat) > 1 {
// and comine array if necessary
				nestedMap[prefix], err = json.Marshal(unflat)
			} else {
				nestedMap[prefix] = unflat[0]
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return
}

// simple O(N) algorithm for removing elements from a slice
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

// convenience function for copying values from one 
// map to anothet
func updateJSONMap(in jsonMap,
	out jsonMap) {
	for key, value := range in {
		out[key] = value
	}
}


