package cmd

import (
	"strings"
	"testing"

	"github.com/XANi/loremipsum"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_jsonHighlight(t *testing.T) {

	json := jsonHighlight(`{"Filter":"string","ID":"string","MaintenanceMode":"string"}`)
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
