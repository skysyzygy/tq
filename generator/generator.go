package main

import (
	"fmt"
	"os"
	"reflect"
	"slices"
	"strings"
	"text/template"

	"github.com/skysyzygy/tq/client"
)

func main() {
	tmpl, err := template.ParseFiles("commands.go.tmpl")
	if err != nil {
		panic(err)
	}

	for _, op := range []string{"Get", "Put", "Post"} {
		outFile, err := os.Create("../cmd/" + strings.ToLower(op) + ".go")
		if err != nil {
			panic(err)
		}

		client := client.New(nil, nil)
		doer, ok := reflect.TypeOf(*client).FieldByName(op)
		if !ok {
			panic(fmt.Errorf("couldn't get client.%v", op))
		}

		// Group commands by the thing they operate on
		commands := make(map[string][]command)
		for i := 0; i < doer.Type.NumMethod(); i++ {
			command := newCommand(doer.Type.Method(i))
			if command.Thing != "" {
				commands[command.Thing] = append(commands[command.Thing], command)
			}
		}

		// Ensure that the first command is the one without a variant
		for _, commands := range commands {
			slices.SortFunc(commands, func(a command, b command) int {
				return strings.Compare(a.Variant, b.Variant)
			})
		}

		data := make(map[string]any)
		data["op"] = op
		data["commands"] = commands
		data["makeAliases"] = makeAliases
		err = tmpl.Execute(outFile, data)
		if err != nil {
			panic(err)
		}
	}
}
