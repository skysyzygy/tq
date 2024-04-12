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
	"io"
	"os"
	"reflect"

	"encoding/json"

	"github.com/go-openapi/runtime"
	"github.com/skysyzygy/tq/auth"
	"github.com/skysyzygy/tq/client"
)

// import "github.com/skysyzygy/tq/cmd"

// Global config options and structures for tq
type tqConfig struct {
	// the Tessitura API client
	client.TessituraServiceWeb

	// Basic auth for requests
	basicAuth func(*runtime.ClientOperation)

	// Bearer token for requests - not yet implemented
	// tokenAuth func(*runtime.ClientOperation)

	// some flags
	verbose bool
	dryRun  bool
}

// Log in the Tessitura client with the given authentication info and cache the login data
func (tq *tqConfig) Login(a auth.Auth) error {

	if validLogin, err := a.Validate(); err != nil || !validLogin {
		return errors.Join(fmt.Errorf("invalid login"), err)
	}
	// Cache the login data
	if basicAuth, err := a.BasicAuth(); err != nil {
		return err
	} else {
		tq.basicAuth = basicAuth
	}
	return nil
}

// Getter function for tq -- simply dispatches to `doGet` based on the type of thing we are querying.
func (tq tqConfig) Get(thing string, query []byte) (res []byte, err error) {

	if tq.verbose {
		fmt.Printf("running Get for %v", thing)
	}

	switch thing {
	case "customers":
		f := tq.TessituraServiceWeb.Get.ConstituentsGetConstituents
		res, err = DoOne(tq, f, query)
		return

	case "emails":
		f := tq.TessituraServiceWeb.Get.ElectronicAddressesGetAll
		res, err = DoOne(tq, f, query)
		return

	case "addresses":
		f := tq.TessituraServiceWeb.Get.AddressesGetAll
		res, err = DoOne(tq, f, query)
		return

	default:
		return nil, fmt.Errorf("no `thing` specified, nothing to do")
	}

}

// Generic for doing (one) thing (get/put/post)
func DoOne[P any, R any, O any, F func(*P, ...O) (R, error)](
	tq tqConfig, function F, query []byte,
) ([]byte, error) {

	params := new(P)
	remainder, err := unmarshallStructWithRemainder(query, params)

	if tq.verbose {
		unmarshallStructReport(params, remainder, os.Stdout)
	}
	if tq.dryRun || err != nil {
		return nil, err
	}

	// Call the function with basic authentication
	obj, err := function(params, interface{}(tq.basicAuth).(O))

	// Marshall the json response
	res, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func main() {
	// cmd.Execute()

}

// Unmarshall into a struct and return the remainder as a map
// Errors if P is not a struct type
func unmarshallStructWithRemainder[P any](query []byte, params *P) (res map[string]any, err error) {
	if reflect.TypeOf(*params).Kind() != reflect.Struct {
		return nil, fmt.Errorf("params must be struct, got %v", reflect.TypeOf(*params).Kind())
	}

	// Unmarshal the query into the given parameter structure
	_err := json.Unmarshal(query, params)
	err = errors.Join(err, _err)

	// Get all the keys of the query for comparison
	_err = json.Unmarshal(query, &res)
	err = errors.Join(err, _err)

	// Remove keys that are in the struct already
	typ := reflect.TypeOf(*params)
	for key := range res {
		if _, ok := typ.FieldByName(key); ok {
			delete(res, key)
		}
	}

	return
}

// Print a message about the different keys in the struct `s` and map `m`
func unmarshallStructReport(s any, m map[string]any, stdout io.WriteCloser) {
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
	// Get the fields that marshmallow didn't map
	mapFields := reflect.ValueOf(m).MapKeys()
	fmt.Fprintf(stdout, "executing a query using fields %v and ignoring fields %v", usedFields, mapFields)
	stdout.Close()
}
