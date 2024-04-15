/*
Copyright © 2024 Sky Syzygy ssyzygy@bam.org

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"encoding/json"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/skysyzygy/tq/auth"
	"github.com/skysyzygy/tq/client"
)

// import "github.com/skysyzygy/tq/cmd"

// Global config options and structures for tq
type tqConfig struct {
	// the Tessitura API client
	*client.TessituraServiceWeb

	// Basic auth for requests
	basicAuth runtime.ClientAuthInfoWriter

	// TODO: Bearer token for requests
	// tokenAuth func(*runtime.ClientOperation)

	// TODO: Logging to file/stdout
	// og io.WriteCloser

	// some flags
	verbose bool
	dryRun  bool
}

// Log in the Tessitura client with the given authentication info and cache the login data
func (tq *tqConfig) Login(a auth.Auth) error {

	// Cache the login data
	if basicAuth, err := a.BasicAuth(); err != nil {
		return err
	} else {
		tq.basicAuth = basicAuth
	}

	host := append(strings.SplitN(a.Hostname, "/", 2), "")
	ignoreCerts, _ := httptransport.TLSClient(httptransport.TLSClientOptions{
		InsecureSkipVerify: true,
	})
	transport := httptransport.NewWithClient(host[0], host[1], []string{"https"}, ignoreCerts)
	transport.DefaultAuthentication = tq.basicAuth
	tq.TessituraServiceWeb = client.New(transport, nil)

	return nil
}

// Generic for doing operations (get/put/post), parallelizing calls to DoOne as needed
// Returns the result, either a single json map if there is only one operation to do,
// or an array of json maps if there are multiple operations.
// Returns the last error
func Do[P any, R any, O any, F func(*P, ...O) (R, error)](
	tq tqConfig, function F, query []byte,
) ([]byte, error) {
	queries := new([]json.RawMessage)
	err := json.Unmarshal(query, queries)
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		// it's not an array... so just call DoOne
		return DoOne(tq, function, query)
	} else if err == nil {
		// loop over queries and call DoOne for each
		// TODO: Parallelize this!
		out := []byte(`[`)
		errs := make([]error, len(*queries))
		var res []byte
		for i, q := range *queries {
			res, errs[i] = DoOne(tq, function, q)
			out = slices.Concat(out, res)
		}
		out = slices.Concat(out, []byte(`]`))
		return out, errs[len(errs)]
	}
	return nil, err // json Unmarshall error
}

// Generic for doing operations (get/put/post), parallelizing calls to DoOne as needed
// The type parameters allow it to work for any swagger-defined function.
// It unmarshals `query` into the appropriate data structure for `function` and marshals
// the result back to json.
// If the initial unmarhshal fails, it tries again with a depth-first traversal of the
// data structure, essentially trying to recast a flat query into the nested structure
// required.
func DoOne[P any, R any, O any, F func(*P, ...O) (R, error)](
	tq tqConfig, function F, query []byte,
) ([]byte, error) {

	params := new(P)
	remainder, err := unmarshallStructWithRemainder(query, params)

	// Is there a better way to determine failure?
	if len(structFields(*params)) == 1 && len(mapFields(remainder)) > 0 {
		remainder, err = unmarshallNestedStructWithRemainder(query, params)
	}

	if tq.verbose {
		structFields(*params)
		mapFields(remainder)
	}
	if tq.dryRun || err != nil {
		return nil, err
	}

	// Call the function
	obj, err := function(params)

	if err != nil {
		return nil, err
	} else {
		// Marshall the json response
		if payload := reflect.ValueOf(obj).Elem().FieldByName("Payload"); !payload.IsZero() {
			return json.Marshal(payload.Interface())
		} else {
			return json.Marshal(obj)
		}
	}
}

func main() {
	// cmd.Execute()

}

// Unmarshall into a struct and return the remainder as a map
// Errors if P is not a struct type
func unmarshallStructWithRemainder(query []byte, params any) (res map[string]any, err error) {
	if reflect.TypeOf(params).Kind() != reflect.Pointer {
		return nil, fmt.Errorf("params must be pointer to struct, got %v", reflect.TypeOf(params).Kind())
	} else if reflect.TypeOf(params).Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("params must be pointer to struct, got pointer to %v", reflect.TypeOf(params).Elem().Kind())
	}

	// Unmarshal the query into the given parameter structure
	_err := json.Unmarshal(query, params)
	err = errors.Join(err, _err)

	// Get all the keys of the query for comparison
	_err = json.Unmarshal(query, &res)
	err = errors.Join(err, _err)

	// Remove keys that are in the struct already
	typ := reflect.TypeOf(params).Elem()
	for key := range res {
		if _, ok := typ.FieldByName(key); ok {
			delete(res, key)
		}
	}

	return
}

// Unmarshall into a nested struct from a flat query and return the remainder as a map
// Errors if P is not a struct type
func unmarshallNestedStructWithRemainder(query []byte, params any) (res map[string]any, err error) {
	if reflect.TypeOf(params).Kind() != reflect.Pointer {
		return nil, fmt.Errorf("params must be pointer to struct, got %v", reflect.TypeOf(params).Kind())
	} else if reflect.TypeOf(params).Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("params must be pointer to struct, got pointer to %v", reflect.TypeOf(params).Elem().Kind())
	}

	v := reflect.ValueOf(params).Elem()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Struct && v.Field(i).IsZero() {
			// Recurse if there's a struct field
			res, err = unmarshallNestedStructWithRemainder(query, v.Field(i).Addr().Interface())
			if err != nil {
				return
			}
			// Update the query with only unmatched fields
			query, _ = json.Marshal(res)
		}
	}

	res, err = unmarshallStructWithRemainder(query, params)

	return
}

// Return the field names from struct `s` that are not zero/nil
func structFields(s any) []string {
	typ := reflect.TypeOf(s)
	val := reflect.ValueOf(s)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	usedFields := make([]string, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		// Get the fields actually assigned in the params struct
		if !val.Field(i).IsZero() {
			usedFields = append(usedFields, typ.Field(i).Name)
		}
	}
	return usedFields
}

// Return the field names from map `m`
func mapFields(m map[string]any) []string {
	fields := make([]string, len(m))
	for field := range m {
		fields = append(fields, field)
	}
	return fields
}
