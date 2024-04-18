package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/spec"
)

// embed the swagger spec
//
//go:embed "tessitura api.15.2.40.json"
var api []byte

// just needed to keep embed in imports for the linter
var _ embed.FS

// Reads the Tessitura swagger spec and extracts documentation info
func Describe[P any, R any, O any, F func(*P, ...O) (*R, error)](
	function F) (short string, long string) {
	// load the swagger spec
	swagger := new(spec.Swagger)
	swagger.UnmarshalJSON(api)
	spec.ExpandSpec(swagger, nil)

	// get function name and split it twice
	funcNameParts := strings.Split(runtime.FuncForPC(
		reflect.ValueOf(function).Pointer()).Name(), "/")
	funcNameParts = strings.Split(
		funcNameParts[len(funcNameParts)-1], ".")

	// get, post, put, delete, head
	funcType := strings.ReplaceAll(funcNameParts[0], "_", "")
	// the function name
	funcName := funcNameParts[2]

	// now find the corresponding operation in the swagger spec!
	var op *spec.Operation
	for _, path := range swagger.Paths.Paths {
		switch funcType {
		case "get":
			op = path.Get
		case "put":
			op = path.Put
		case "post":
			op = path.Post
		case "head":
			op = path.Head
		case "delete":
			op = path.Delete
		}

		if op != nil &&
			// funcName is often appended in golang by a suffix
			// (oftem `-fm` to make the name unique)
			// https://stackoverflow.com/questions/32925344/why-is-there-a-fm-suffix-when-getting-a-functions-name-in-go
			strings.Contains(funcName,
				// the swagger doc has an extra "_" in the name
				strings.ReplaceAll(op.ID, "_", "")) {
			break
		}
	}

	short = strings.ReplaceAll(strings.ReplaceAll(op.ID, "_", ""), funcType, "")
	long = op.Summary

	spew := spew.NewDefaultConfig()
	spew.MaxDepth = 2

	params := make(map[string]any)

	for _, param := range op.Parameters {
		name := param.Name
		if param.Required {
			name += "*"
		}
		if param.Schema != nil && param.Schema.SchemaProps.Properties != nil {
			params_name := make(map[string]string)
			for key, prop := range param.Schema.SchemaProps.Properties {
				params_name[key] = prop.Type[0]
			}
			params[name] = params_name
		} else {
			params[name] = param.Type
		}
	}
	res, _ := json.MarshalIndent(params, ">", "   ")
	fmt.Println(string(res))
	return
}
