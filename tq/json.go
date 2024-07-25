package tq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

// String map of json raw strings
type jsonMap map[string]json.RawMessage
type records [][]string

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
func FlattenJSONMap(in json.RawMessage, prefix string) (flatMap jsonMap, err error) {
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
			flatMapPart, err = FlattenJSONMap(value, prefix+key)
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
			flatMapPart, err := FlattenJSONMap(value, prefix+"["+strconv.Itoa(i)+"]")
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
	slices.Sort(prefixes)
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

// unmarshal json.RawMessage into a slice of jsonMaps
func jsonToJSONMaps(in json.RawMessage) (out []jsonMap, err error) {
	in = bytes.TrimSpace(in)
	if len(in) == 0 {
		return nil, errors.New("input JSON has length 0")
	}
	if in[0] == '[' {
		err = json.Unmarshal(in, &out)
		return
	} else {
		out = make([]jsonMap, 1)
		err = json.Unmarshal(in, &out[0])
		return
	}
}

// Call flattenJSONMap for a json.RawMessage and return a slice of jsonMaps
func flattenJSONMaps(in json.RawMessage) (out []jsonMap, err error) {
	var inArr []json.RawMessage
	in = bytes.TrimSpace(in)
	if len(in) == 0 {
		return nil, errors.New("input JSON has length 0")
	}
	if in[0] == '[' {
		err = json.Unmarshal(in, &inArr)
		if err != nil {
			return nil, err
		}
		out = make([]jsonMap, len(inArr))
		for i, j := range inArr {
			out[i], err = FlattenJSONMap(j, "")
			if err != nil {
				return nil, err
			}
		}
	} else {
		out = make([]jsonMap, 1)
		out[0], err = FlattenJSONMap(in, "")
		if err != nil {
			return nil, err
		}
	}
	return
}

// Call unflattenJSONMap for a slice of jsonMaps and return a json.RawMessage byte slice
func unflattenJSONMaps(in []jsonMap) (out json.RawMessage, err error) {
	outArr := make([]json.RawMessage, len(in))
	for i, j := range in {
		outArr[i], err = unflattenJSONMap(j)
		if err != nil {
			return nil, err
		}
	}
	return json.Marshal(outArr)
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

// convenience function for copying values from one
// map to anothet
func updateJSONMap(in jsonMap,
	out jsonMap) {
	for key, value := range in {
		out[key] = value
	}
}
