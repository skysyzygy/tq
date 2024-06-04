package cmd

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/skysyzygy/tq/tq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func testCommand(t *testing.T, command *cobra.Command, flagName string) {
	use := command.Use
	r, w, _ := os.Pipe()

	if flagName != "" {
		flag := command.Flag(flagName)
		if flag != nil {
			flag.Value.Set("true")
			defer flag.Value.Set("false")
			use = flag.Usage
		}
	}
	input := regexp.MustCompile(`\{.+\}$`).FindString(
		strings.ReplaceAll(use, ",...", ""))

	viper.Set("login", authString)
	_tq = tq.New(nil, false, false)
	_tq.SetInput(r)

	w.Write([]byte(input))
	w.Close()
	err := command.Execute()

	assert.NoError(t, err)
	// Note: need to test output better
	assert.NotEmpty(t, string(_tq.GetOutput()))

}
