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
	basicAuth runtime.ClientAuthInfoWriter

	// Bearer token for requests
	tokenAuth runtime.ClientAuthInfoWriter

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

// Getter function for tq -- handles calling gotess API methods, translating incoming requests to filtersets and parallelizing when necessary
func (tq *tqConfig) Get(thing string, query map[string]any) (res any, err error) {
	if tq.verbose {
		fmt.Printf("running Get for %v", thing)
	}
	switch thing {
	case "customers":

	case "emails":

	default:
		return nil, fmt.Errorf("no `thing` specified, nothing to do")
	}

	if tq.verbose {
		fields := make([]string, 0)
		fmt.Printf("executing a query of %v using fields: %v", thing, fields)
	}

	if tq.dryRun {
		return
	}

	return
}

func main() {
	// cmd.Execute()
	tq := tqConfig{}
	tq.Login(auth.Auth{})

}
