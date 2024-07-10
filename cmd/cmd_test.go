package cmd

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/skysyzygy/tq/tq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func testCommand(t *testing.T, command *cobra.Command, flagName string) {
	use := ansi.Strip(command.Example)

	r, w, _ := os.Pipe()
	command.SetIn(r)

	if flagName != "" {
		flag := command.Flag(flagName)
		if flag != nil {
			flag.Value.Set("true")
			defer flag.Value.Set("false")
			use = ansi.Strip(flag.Usage)
		}
	}
	input := regexp.MustCompile(`\{.+\}$`).FindString(
		strings.ReplaceAll(use, ",...", ""))

	viper.Set("login", authString)
	_tq = tq.New(nil, false, false)
	initTq(command, nil)

	w.Write([]byte(input))
	w.Close()
	err := command.RunE(command, nil)

	assert.NoError(t, err)
	// Note: need to test output better
	assert.NotEmpty(t, string(_tq.GetOutput()))

}
