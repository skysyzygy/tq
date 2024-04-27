package main

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_execTemplate(t *testing.T) {
	template := "Hello {{ .world }}"
	inFile := path.Join(os.TempDir(), "generator_test.tmpl")
	outFile := path.Join(os.TempDir(), "execTemplate.txt")

	err := os.WriteFile(inFile, []byte(template), 0)
	assert.NoError(t, err)
	defer os.Remove(inFile)

	data := map[string]string{"world": "Yourself!"}
	assert.NoFileExists(t, outFile)

	// Execute the template
	err = execTemplate(inFile, outFile, data)
	assert.FileExists(t, outFile)
	assert.NoError(t, err)

	// Read the output
	contents, _ := os.ReadFile(outFile)
	assert.Equal(t, "Hello Yourself!", string(contents))

	// Clean up
	os.Remove(outFile)
	assert.NoFileExists(t, outFile)
}

func Test_getDataForOperation(t *testing.T) {
	assert.Panics(t, func() { getDataForOperation("NotAnOp") })

	get := getDataForOperation("Get")
	assert.Equal(t, get["op"], "Get")
	commands, ok := get["commands"].(map[string][]command)
	assert.True(t, ok, "get[`commands`] is a map of slices of `command`s")
	assert.Greater(t, len(commands), 100)

	ops := make(map[string]any)
	for _, op := range []string{"Get", "Put", "Post"} {
		ops[op] = getDataForOperation(op)
	}

	commandHash := make(map[string]bool)
	// Check that there are no matching verb/thing/variant combos
	// and that there is alignment between the data map and command structs
	for op, data := range ops {
		things := data.(map[string]any)["commands"].(map[string][]command)
		for thing, commands := range things {
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
