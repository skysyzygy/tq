package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/go-openapi/spec"
)

// embed the swagger spec
//
//go:embed "tessitura api.15.2.40.json"
var api []byte

// just needed to keep embed in imports for the linter
var _ embed.FS

var swagger *spec.Swagger

// load the swagger spec
func init() {
	swagger = new(spec.Swagger)
	swagger.UnmarshalJSON(api)
	spec.ExpandSpec(swagger, nil)
}

// Reads the Tessitura swagger spec and extracts documentation info
func Describe(funcName string) (short string, long string) {

	// now find the corresponding operation in the swagger spec!
	for url, path := range swagger.Paths.Paths {
		for _, op := range []*spec.Operation{
			path.Get, path.Put, path.Post, path.Delete, path.Head,
		} {
			// funcName is often appended in golang by a suffix
			// (oftem `-fm` to make the name unique)
			// https://stackoverflow.com/questions/32925344/why-is-there-a-fm-suffix-when-getting-a-functions-name-in-go
			if op != nil && funcName ==
				// the swagger doc has an extra "_" in the name
				strings.ReplaceAll(op.ID, "_", "") {
				url := strings.Split(url, "/")
				short = url[1]
				if len(url) > 2 {
					short = url[2]
				}
				long = op.Summary
				for _, param := range op.OperationProps.Parameters {
					paramReq := ""
					if param.Required {
						paramReq = "(*)"
					}
					paramType := param.Type
					if paramType == "" {
						paramType = param.Schema.Type[0]
					}
					paramName := strings.ToUpper(string(param.Name[0])) +
						param.Name[1:]
					long += fmt.Sprintf("\n %v%v: %v\t%v", paramName, paramReq,
						paramType, param.Description)
				}

				return
			}
		}
	}

	return
}

// Instantiate a struct type and fill it with descriptions up to depth (starting at 0)
func instantiateStructType(t reflect.Type, depth int) any {
	var v reflect.Value

	// instantiate the root element
	if t.Kind() == reflect.Pointer &&
		t.Elem().Kind() == reflect.Struct {
		// v = *new(*t)
		v = reflect.New(t.Elem()).Elem()
	} else if t.Kind() == reflect.Struct {
		// v = *new(t)
		v = reflect.New(t).Elem()
	} else {
		return reflect.New(t).Interface()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.CanSet() {
			if field.Kind() == reflect.Pointer {
				field.Set(reflect.New(field.Type().Elem()))
				field = field.Elem()
			}
			switch field.Kind() {
			// bools, ints and floats are already initialized
			case reflect.String:
				field.SetString("string")
			case reflect.Struct:
				if depth > 0 {
					field.Set(reflect.ValueOf(instantiateStructType(field.Type(), depth-1)))
				} else {
					// unset pointer
					v.Field(i).SetZero()
				}
			}
		}
	}

	if t.Kind() == reflect.Pointer {
		return v.Addr().Interface()
	}
	return v.Interface()
}

func usage(method reflect.Value) []byte {
	params := reflect.ValueOf(instantiateStructType(method.Type().In(1), 1))
	if params.Kind() == reflect.Pointer {
		params = params.Elem()
	}
	if params.Kind() != reflect.Struct {
		// Oops! Don't know what to do =)
		return nil
	}

	paramsMap := make(map[string]any)
	removeFields := []string{"timeout", "Context", "HTTPClient"}
	for i := 0; i < params.NumField(); i++ {
		fieldName := params.Type().Field(i).Name
		if !slices.Contains(removeFields, fieldName) {
			paramsMap[fieldName] = params.Field(i).Interface()
		}
	}

	usage, err := json.Marshal(paramsMap)
	if err != nil {
		panic(err)
	}
	usage = []byte(strings.ReplaceAll(string(usage), "null", "[object]"))
	return usage
}

// Fills a command struct with info about a method
func newCommand(method reflect.Method) command {
	short, long := Describe(method.Name)
	stopWords := []string{"Get", "Update", "Create", "Delete", short}
	variant := method.Name
	verb := ""
	for _, stopWord := range stopWords {
		if strings.Contains(variant, stopWord) && verb == "" {
			verb = stopWord
		}
		variant = strings.ReplaceAll(variant, stopWord, "")
	}
	return command{
		verb:    verb,
		thing:   short,
		variant: variant,
		long:    long,
		usage:   string(usage(method.Func)),
		aliases: []string{short},
	}

}
