package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/skysyzygy/tq/tq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var test_cmd = &cobra.Command{
	PreRunE: tqInit,
	RunE: func(cmd *cobra.Command, args []string) error {
		_tq.Log.Info(args[0])
		return fmt.Errorf(args[0])
	},
}

// test that log and file options are getting set for cobra commands by tqInit
func Test_tqInit(t *testing.T) {
	test_json := []byte(`{"some":"json"}`)

	os.WriteFile("test.json", test_json, 0644)
	defer os.Remove("test.json")

	jsonFile = "test.json"
	defer func() { jsonFile = "" }()

	logFile = os.TempDir() + string(os.PathSeparator) + "test.log"
	defer func() { logFile = "" }()

	viper.Set("login", authString)

	// test that input file is getting read
	err := tqInit(test_cmd, nil)
	assert.NoError(t, err)
	tq.CaptureOutput(func() {
		err = test_cmd.Execute()
	})
	assert.ErrorContains(t, err, string(test_json))

	// test that log file is getting written to
	assert.FileExists(t, logFile)
	log, _ := os.ReadFile(logFile)
	assert.Contains(t, string(log), "\\\"some\\\":\\\"json\\\"")
}

func Test_tqInit_Errors(t *testing.T) {
	jsonFile = "test.json"
	viper.Set("login", authString)
	var err error

	// test that absent input file throws an error
	tq.CaptureOutput(func() {
		err = tqInit(test_cmd, nil)
	})
	assert.ErrorContains(t, err, "cannot open input file for reading")
	jsonFile = ""

	// test that unwriteable log file throws an error
	logFile = "not_a_dir/test.log"
	tq.CaptureOutput(func() {
		err = tqInit(test_cmd, nil)
	})
	assert.ErrorContains(t, err, "cannot open log file for appending")
	logFile = ""

}
