package tq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

// String map of json raw strings
type jsonMap map[string]json.RawMessage
type csv [][]string

func flattenJSONError(key string, v any) error {
	if err, ok := v.(error); ok {
		return errors.Join(fmt.Errorf("couldn't flatten key %s", key), err)
	} else {
		panic(v)
	}
}

func unflattenJSONError(key string, v any) error {
	if err, ok := v.(error); ok {
		return errors.Join(fmt.Errorf("couldn't unflatten key %s", key), err)
	} else {
		panic(v)
	}
}

// Flattens JSON input with possibly nested values to an unnested
// (flat) jsonMap such that each key is the JSONPath location of the value
// in the input object.
//
// Assumptions:
// There are no duplicate keys
// No keys contain the characters .[]
func flattenJSONMap(in json.RawMessage, prefix string) (flatMap jsonMap, err error) {
	var (
		key   string
		value json.RawMessage
	)
	defer func() {
		if r := recover(); r != nil {
			err = flattenJSONError(prefix+key, r)
		}
	}()

	var flatMapPart jsonMap
	flatMap = make(jsonMap)

	in = bytes.TrimSpace(in)
	if len(in) == 0 {
		return nil, nil
	}
	switch string(in[0]) {
	case "{":
		// make json a map
		nestedMapPart := make(jsonMap)
		err = json.Unmarshal(in, &nestedMapPart)
		if err != nil {
			panic(err)
		}
		if prefix != "" {
			prefix = prefix + "."
		}
		// recurse
		for key, value = range nestedMapPart {
			flatMapPart, err = flattenJSONMap(value, prefix+key)
			if err != nil {
				panic(err)
			}
			updateJSONMap(flatMapPart, flatMap)
		}
	case "[":
		// split json array
		array := make([]json.RawMessage, 0, strings.Count(string(in), ","))
		err = json.Unmarshal(in, &array)
		if err != nil {
			panic(err)
		}
		for i, value := range array {
			// recurse
			flatMapPart, err := flattenJSONMap(value, prefix+"["+strconv.Itoa(i)+"]")
			if err != nil {
				panic(err)
			}
			updateJSONMap(flatMapPart, flatMap)
		}
	default:
		// handle atomic values
		flatMap[prefix] = in
	}

	return
}

// Unflattens a jsonMap by parsing each key as the JSONPath location of the
// value in the final jsonMap.
//
// Assumptions:
// There are no duplicate keys
func unflattenJSONMap(flatMap jsonMap) (out json.RawMessage, err error) {
	var (
		prefix string
		value  json.RawMessage
	)
	nestedMap := make(map[string]jsonMap)
	outMap := make(jsonMap)

	defer func() {
		if r := recover(); r != nil {
			err = unflattenJSONError(prefix, r)
		}
	}()

	for key := range flatMap {
		value = flatMap[key]
		prefix = ""
		if sep := strings.IndexAny(key, ".[]"); sep != -1 {
			prefix = key[0:sep]
			key = key[sep+1:]
			// keep removing control characters
			for len(key) > 0 && strings.ContainsAny(key[0:1], ".[]") {
				key = key[1:]
			}
		} else {
			// don't recurse on bare keys
			outMap[key] = value
			continue
		}
		if nestedMap[prefix] == nil {
			nestedMap[prefix] = make(jsonMap)
		}
		nestedMap[prefix][key] = value
	}

	prefixes := maps.Keys(nestedMap)
	if len(prefixes) > 0 {
		prefix = prefixes[0]
		// is array
		if _, err := strconv.Atoi(prefix); err == nil {
			maxIndex, err := sliceMax(prefixes)
			if err != nil {
				panic(err)
			}
			outArr := make([]json.RawMessage, maxIndex+1)
			for _, prefix = range prefixes {
				i, _ := strconv.Atoi(prefix)
				keys := maps.Keys(nestedMap[prefix])
				// unkeyed values in arrays are atomic and don't need recursion
				if len(keys) == 1 && keys[0] == "" {
					outArr[i] = nestedMap[prefix][""]
				} else {
					outArr[i], err = unflattenJSONMap(nestedMap[prefix])
				}
				if err != nil {
					panic(err)
				}
			}
			return json.Marshal(outArr)
		} else {
			for _, prefix = range prefixes {
				outMap[prefix], err = unflattenJSONMap(nestedMap[prefix])
				if err != nil {
					panic(err)
				}
			}
		}
	}

	return json.Marshal(outMap)
}

// Convert a slice of jsonMap to a slice of csv records plus a header row ([]string)
func jsonMapsToCsv(in []jsonMap) (out csv) {
	keys := make(map[string]bool)
	out = make(csv, len(in)+1)

	// gather all keys
	for _, row := range in {
		for key := range row {
			keys[key] = true
		}
	}

	// fill in the csv
	i := 0
	out[0] = make([]string, len(keys))
	for key := range keys {
		out[0][i] = key
		for j, row := range in {
			if out[j+1] == nil {
				out[j+1] = make([]string, len(keys))
			}
			out[j+1][i] = string(row[key])
		}
		i++
	}
	return
}

// Convert a slice of csv records plus a header row ([]string) to a slice of jsonMap
func jsonMapsFromCsv(in csv) (out []jsonMap, err error) {
	if len(in) == 0 {
		return nil, errors.New("csv has no rows")
	}

	keys := in[0]
	out = make([]jsonMap, len(in)-1)

	for i := 1; i < len(in); i++ {
		out[i-1] = make(jsonMap)
		for j, key := range keys {
			out[i-1][key] = []byte(in[i][j])
		}
	}

	return
}

func sliceMax(s []string) (int, error) {
	maxIndex := 0
	for _, p := range s {
		i, err := strconv.Atoi(p)
		if i > maxIndex {
			maxIndex = i
		}
		if err != nil {
			return 0, err
		}
	}
	return maxIndex, nil
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
