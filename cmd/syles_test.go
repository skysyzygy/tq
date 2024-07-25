package cmd

import (
	"strings"
	"syscall"
	"testing"

	"github.com/XANi/loremipsum"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type termTester struct{}

func (t termTester) IsTerminal(int) bool { return isTerminal }

var isTerminal bool

func Test_jsonHighlight(t *testing.T) {
	stdout := syscall.Stdout
	defer func() {
		syscall.Stdout = stdout
		highlight = false
		noHighlight = false
		terminal = xTerm{}
	}()

	terminal = termTester{}
	isTerminal = true
	// highlights by default
	json := jsonStyle(`{"Filter":"string","ID":"string","MaintenanceMode":"string"}`, false)
	// contains ANSI escape sequence
	assert.Contains(t, json, "\033")

	// but not when noHighlight is set
	noHighlight = true
	json = jsonStyle(`{"Filter":"string","ID":"string","MaintenanceMode":"string"}`, false)
	// doesn't contain ANSI escape sequence
	assert.NotContains(t, json, "\033")

	// or when output is not to terminal
	isTerminal = false
	json = jsonStyle(`{"Filter":"string","ID":"string","MaintenanceMode":"string"}`, false)
	// doesn't contain ANSI escape sequence
	assert.NotContains(t, json, "\033")

	// unless highlight is set
	highlight = true
	json = jsonStyle(`{"Filter":"string","ID":"string","MaintenanceMode":"string"}`, false)
	// contains ANSI escape sequence
	assert.Contains(t, json, "\033")

}

func Test_flagUsagesWrapped(t *testing.T) {

	generator := loremipsum.New()
	words := generator.Paragraph()

	cmd := cobra.Command{}
	cmd.Flags().Bool("test", false, words)

	for i := 16; i <= 255; i++ {
		wrapped := flagUsagesWrapped(i, cmd.Flags())
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
	ew := exampleWrapped(1024, example)
	assert.Contains(t, ew, `"nested"`)

	*flatHelp = true
	defer func() { *flatHelp = false }()
	ew = exampleWrapped(1024, example)
	assert.Contains(t, ew, `"nested[0]"`)

	ew = exampleWrapped(8, example)
	assert.Contains(t, ew, "Filter\n")
}
