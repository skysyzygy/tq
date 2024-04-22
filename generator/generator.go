package main

import (
	"os"
	"reflect"
	"text/template"

	"github.com/skysyzygy/tq/client"
)

type command struct {
	verb    string
	thing   string
	variant string
	long    string
	usage   string
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
