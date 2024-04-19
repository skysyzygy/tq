package main

import (
	"encoding/json"
	"os"
	"reflect"
	"text/template"

	"github.com/skysyzygy/tq/client"
)

type command struct {
	use     string
	short   string
	long    string
	name    string
	aliases []string
}

func main() {
	tmpl, err := template.ParseFiles("commands.go.tmpl")
	if err != nil {
		panic(err)
	}

	outFile, _ := os.OpenFile("get.go", os.O_RDWR, 0644)

	client := client.New(nil, nil)
	get := reflect.TypeOf(client.Put)
	commands := make([]command, get.NumMethod())
	for i := 0; i < get.NumMethod(); i++ {
		commands[i] = newCommand(get.Method(i))
	}
	tmpl.Execute(outFile, commands)
}

// Instantiate a struct type and fill it with descriptions
func instantiateStructType(t reflect.Type) any {
	var v reflect.Value
	// instantiate the root element
	if t.Kind() == reflect.Pointer {
		// v = *new(*t)
		v = reflect.New(t.Elem()).Elem()
	} else {
		// v = *new(t)
		v = reflect.New(t).Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.CanSet() {
			switch field.Kind() {
			// bools, ints and floats are already initialized
			case reflect.String:
				field.SetString("string")
			case reflect.Pointer, reflect.Struct:
				// recurse!
				field.Set(reflect.ValueOf(instantiateStructType(field.Type())))
			}
		}
	}

	if t.Kind() == reflect.Pointer {
		return v.Addr().Interface()
	}
	return v.Interface()
}

// Fills a command struct with info about a method
func newCommand(method reflect.Method) command {

	name := method.Name
	short, long := DescribeFunc(name)
	params := instantiateStructType(method.Func.Type().In(1).Elem())
	use, _ := json.Marshal(params)
	return command{
		name:    name,
		short:   short,
		long:    long,
		use:     string(use),
		aliases: []string{name},
	}

}
