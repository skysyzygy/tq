package main

import (
	"embed"
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"regexp"
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

func DescribeFunc(funcName string) (short string, long string) {
	paths := []string{"g_e_t", "p_u_t", "p_o_s_t", "d_e_l_e_t_e", "h_e_a_d"}
	for _, path := range paths {
		re := regexp.MustCompile(`\s*` + funcName + `\s*`)
		pkgs, _ := parser.ParseDir(token.NewFileSet(), "../client/"+path, nil, parser.ParseComments)
		for _, pkg := range pkgs {
			doc := doc.New(pkg, "../client", 0)
			for _, typ := range doc.Types {
				if typ.Name == "Client" {
					for _, fun := range typ.Methods {
						if fun.Name == funcName {
							short = fun.Name
							long = re.ReplaceAllString(fun.Doc, "")
							return
						}
					}

				}
			}
		}

	}
	return
}

// Reads the Tessitura swagger spec and extracts documentation info
func Describe(funcName string) (short string, long string) {

	// now find the corresponding operation in the swagger spec!
	for _, path := range swagger.Paths.Paths {
		for _, op := range []*spec.Operation{
			path.Get, path.Put, path.Post, path.Delete, path.Head,
		} {
			// funcName is often appended in golang by a suffix
			// (oftem `-fm` to make the name unique)
			// https://stackoverflow.com/questions/32925344/why-is-there-a-fm-suffix-when-getting-a-functions-name-in-go
			if op != nil && funcName ==
				// the swagger doc has an extra "_" in the name
				strings.ReplaceAll(op.ID, "_", "") {
				short = funcName
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
