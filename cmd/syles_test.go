package cmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/XANi/loremipsum"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_jsonHighlight(t *testing.T) {

	json := jsonHighlight(`{"Filter":"string","ID":"string","MaintenanceMode":"string","nested":["bird","squirrel","rat"]}`, false)
	// contains ANSI escape sequence
	assert.Contains(t, json, "\033")
	assert.NotContains(t, json, "nested[0]")

	json = jsonHighlight(`{"Filter":"string","ID":"string","MaintenanceMode":"string","nested":["bird","squirrel","rat"]}`, true)
	fmt.Println(json)
	assert.Contains(t, json, "\033")
	assert.Contains(t, json, "nested[0]")
}

func Test_flagUsagesWrapped(t *testing.T) {

	generator := loremipsum.New()
	words := generator.Paragraph()

	cmd := cobra.Command{}
	cmd.Flags().Bool("test", false, words)

	for i := 16; i <= 255; i++ {
		wrapped := flagUsagesWrapped(i, false, cmd.Flags())
		for _, line := range strings.Split(wrapped, "\n") {
			if len(line) > 1 {
				assert.Equal(t, "  ", line[0:2])
			}
			assert.LessOrEqual(t, len(line), i)
		}
	}

}

func Test_exampleWrapped(t *testing.T) {
	example := `{"Filter":"string","ID":"string","MaintenanceMode":"string","nested":["bird","squirrel","rat"]}`
	ew := exampleWrapped(1024, false, example)
	assert.Contains(t, ew, `"nested"`)

	ew = exampleWrapped(1024, true, example)
	assert.Contains(t, ew, `"nested[0]"`)

	ew = exampleWrapped(8, true, example)
	assert.Contains(t, ew, "Filter\n")
}
