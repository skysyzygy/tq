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

	"encoding/json"

	"github.com/go-openapi/runtime"
	"github.com/perimeterx/marshmallow"
	"github.com/skysyzygy/tq/auth"
	"github.com/skysyzygy/tq/client"
	"github.com/skysyzygy/tq/client/g_e_t"
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
		p := g_e_t.ConstituentsGetConstituentsParams{}
		f := tq.TessituraServiceWeb.Get.ConstituentsGetConstituents
		res, err = doGet(tq, p, f, query)
		return

	case "emails":
		p := g_e_t.ElectronicAddressesGetAllParams{}
		f := tq.TessituraServiceWeb.Get.ElectronicAddressesGetAll
		res, err = doGet(tq, p, f, query)
		return

	case "addresses":
		p := g_e_t.AddressesGetAllParams{}
		f := tq.TessituraServiceWeb.Get.AddressesGetAll
		res, err = doGet(tq, p, f, query)
		return

	default:
		return nil, fmt.Errorf("no `thing` specified, nothing to do")
	}

}

// Generic for getting things
func doGet[P any, R any, F func(*P, ...g_e_t.ClientOption) (R, error)](
	tq tqConfig, params P, function F, query []byte,
) (res []byte, err error) {

	// Try to unmarshal the query into the given parameter structure
	rest, _err := marshmallow.Unmarshal(query, &params, marshmallow.WithExcludeKnownFieldsFromMap(true))
	err = errors.Join(err, _err)

	if tq.verbose {
		typ := reflect.TypeOf(params)
		val := reflect.ValueOf(params)
		fields := make([]string, typ.NumField())
		usedFields := make([]string, 0, typ.NumField())
		for i := range fields {
			fields[i] = typ.Field(i).Name
			// Get the fields actually assigned in the params struct
			if val.Field(i).String() != "" {
				usedFields[len(usedFields)] = typ.Field(i).Name
			}
		}
		// Get the fields that marshmallow didn't map
		restFields := reflect.ValueOf(rest).MapKeys()
		fmt.Printf("executing a query using fields %v and ignoring fields %v", usedFields, restFields)
	}
	if tq.dryRun {
		return nil, err
	}

	obj, _err := function(&params, tq.basicAuth)
	err = errors.Join(err, _err)

	res, _err = json.Marshal(obj)
	err = errors.Join(err, _err)

	return res, err
}

func main() {
	// cmd.Execute()
	tq := tqConfig{}
	tq.Login(auth.Auth{})

}
