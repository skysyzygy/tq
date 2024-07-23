package tq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

// String map of json raw strings
type jsonMap map[string]json.RawMessage

func flattenJSONError(key string, v any) error {
	if err, ok := v.(error); ok {
		return errors.Join(errors.New(fmt.Sprintf("couldn't flatten key %s", key)), err)
	} else {
		panic(v)
	}
}

func unflattenJSONError(key string, v any) error {
	if err, ok := v.(error); ok {
		return errors.Join(errors.New(fmt.Sprintf("couldn't unflatten key %s", key)), err)
	} else {
		panic(v)
	}
}

func flattenJSON(in json.RawMessage, prefix string) (out jsonMap, err error) {
	inMap := make(jsonMap)
	out = make(jsonMap)
	err = json.Unmarshal(in, &inMap)
	if err != nil {
		out[prefix] = in
		err = nil
		return
	}
	return flattenJSONMap(inMap, prefix)
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
	var (
		key   string
		value json.RawMessage
	)
	defer func() {
		if r := recover(); r != nil {
			err = flattenJSONError(key, r)
		}
	}()
	if prefix != "" {
		prefix = prefix + "."
	}
	for key, value = range nestedMap {
		flatMapPart := make(jsonMap)
		value = bytes.TrimSpace(value)
		switch string(value[0]) {
		case "{":
			// make json a map
			nestedMapPart := make(jsonMap)
			err = json.Unmarshal(value, &nestedMapPart)
			if err != nil {
				panic(err)
			}
			// recurse
			flatMapPart, err = flattenJSONMap(nestedMapPart, prefix+key)
			if err != nil {
				panic(err)
			}
		case "[":
			// split json array
			array := make([]json.RawMessage, 0, strings.Count(string(value), ","))
			err = json.Unmarshal(value, &array)
			if err != nil {
				panic(err)
			}
			for i, value := range array {
				// recurse back to flattenJSON to allow atomic elements
				flatMapPartPart, err := flattenJSON(value, prefix+key+"["+strconv.Itoa(i)+"]")
				if err != nil {
					panic(err)
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

	var (
		key   string
		index int
		value json.RawMessage
	)
	nestedMap = make(jsonMap)
	nestedMapPart := make(map[string][]jsonMap)

	defer func() {
		if r := recover(); r != nil {
			err = unflattenJSONError(key, r)
		}
	}()
	keys := maps.Keys(flatMap)
	sort.Slice(keys, func(i, j int) bool {
		return i < j
	})
	for _, key = range keys {
		// bare key, nothing to do
		if sep := strings.IndexAny(key, ".[]"); sep == -1 {
			nestedMap[key] = flatMap[key]
		} else {
			prefix := key[0 : sep+1]
			value = flatMap[key]
			index = 0
			key = strings.TrimPrefix(key, prefix)
			if i, _key, found := strings.Cut(key, "]."); found && !strings.Contains(i, ".") {
				index, err = strconv.Atoi(i)
				key = _key
				if err != nil {
					panic(err)
				}
			}
			flatMapPart = append(flatMapPart, make(jsonMap))
			flatMapPart[index][key] = value
		}
		// len(keys) is an upper bound
		for len(keys) > 0 {
			for i := range flatMapPart {
				// recurse
				_unflat, err := unflattenJSONMap(flatMapPart[i])
				if err != nil {
					panic(err)
				}
				// combine into json message
				unflat[i], err = json.Marshal(_unflat)
				if err != nil {
					panic(err)
				}
			}
			if len(unflat) > 1 {
				// and comine array if necessary
				nestedMap[prefix[0:len(prefix)-1]], err = json.Marshal(unflat)
			} else {
				nestedMap[prefix[0:len(prefix)-1]] = unflat[0]
			}
		}
	}
	if err != nil {
		panic(err)
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
