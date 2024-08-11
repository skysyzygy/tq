package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func Test_generate(t *testing.T) {
	defer func() {
		os.Remove("templ.test")
	}()
	data := make(map[string]any)
	data["commands"] = getCommandData()
	data["op"] = "Get"
	generate("generator_test.tmpl", "templ.test", data)
	assert.FileExists(t, "templ.test")
	get, _ := os.ReadFile("templ.test")
	getLines := strings.Split(string(get), "\n")

	assert.Equal(t, "Op: Get", getLines[0])
	assert.Greater(t, len(getLines), 100)
	assert.Contains(t, getLines, "Command: Constituents")
	assert.Contains(t, getLines[1:], "  - Op: Get")
	assert.Contains(t, getLines[1:], "  - Op: Post")
	assert.Contains(t, getLines[1:], "  - Op: Put")
}

func Test_getCommandData(t *testing.T) {

	cmdData := getCommandData()

	assert.Greater(t, len(cmdData), 100)
	assert.Contains(t, maps.Keys(cmdData["Addresses"]), "Get")
	assert.Contains(t, maps.Keys(cmdData["Addresses"]), "Put")
	assert.Contains(t, maps.Keys(cmdData["Addresses"]), "Post")

	commandHash := make(map[string]bool)
	// Check that there are no matching verb/thing/variant combos
	// and that there is alignment between the data map and command structs
	for thing, data := range cmdData {
		for op, commands := range data {
			for i, command := range commands {
				//assert.Equal(t, op, command.Verb) -- verbs are sometimes not the same as the op (e.g. "Create", "Update")
				assert.Equal(t, thing, command.Thing)
				assert.False(t, commandHash[strings.Join([]string{op, thing, command.Variant}, "__")])
				commandHash[strings.Join([]string{op, thing, command.Variant}, "__")] = true

				// ensure that only the first command in a slice can be without variant
				if i != 0 {
					assert.NotEqual(t, "", command.Variant)
				}

			}
		}
	}

}
