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
	for _, op := range []string{"Get", "Put", "Post"} {
		if err := execTemplate("commands.go.tmpl", "../cmd/"+strings.ToLower(op)+".go", getDataForOperation(op)); err != nil {
			panic(err)
		}
	}
}

// Run the template in `inFilename` using `data` and save as `outFilename`
func execTemplate(inFilename string, outFilename string, data any) error {
	tmpl, err := template.ParseFiles(inFilename)
	if err != nil {
		return err
	}
	outFile, err := os.Create(outFilename)

	if err != nil {
		return err
	}
	err = tmpl.Execute(outFile, data)
	if err != nil {
		return err
	}
	return outFile.Close()
}

// Build data about entities that can be used with `operation` (i.e. "Get", "Post", "Put")
func getDataForOperation(operation string) (data map[string]any) {
	client := client.New(nil, nil)
	doer, ok := reflect.TypeOf(*client).FieldByName(operation)
	if !ok {
		panic(fmt.Errorf("couldn't get client.%v", operation))
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

	data = make(map[string]any)
	data["op"] = operation
	data["commands"] = commands
	data["makeAliases"] = makeAliases
	return
}
