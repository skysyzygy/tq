package cmd

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/skysyzygy/tq/tq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var test_cmd = &cobra.Command{
	PreRunE: initTq,
	RunE: func(cmd *cobra.Command, args []string) error {
		input, _ := _tq.ReadInput()
		_tq.Log.Info(string(input))
		return fmt.Errorf(string(input))
	},
}

// test that initLog sets up logging by creating a tq instance
func Test_initLog(t *testing.T) {
	logFile = os.TempDir() + string(os.PathSeparator) + "test.log"
	defer func() { logFile = "" }()

	initLog()
	_tq.Log.Info("Starting log!")

	// test that log file is getting written to
	assert.FileExists(t, logFile)
	log, _ := os.ReadFile(logFile)
	assert.Contains(t, string(log), "Starting log!")

	// test that unwriteable log file throws an error
	logFile = "not_a_dir/test.log"
	_, err := tq.CaptureOutput(func() {
		initLog()
	})
	assert.Regexp(t, "cannot open log file .+ for appending", string(err))
	logFile = ""
}

// test that log and file options are getting set for cobra commands by tqInit
func Test_tqInit(t *testing.T) {
	test_json := []byte(`{"some":"json"}`)

	os.WriteFile("test.json", test_json, 0644)
	defer os.Remove("test.json")

	inFile = "test.json"
	defer func() { inFile = "" }()

	viper.Set("login", authString)

	logFile = os.TempDir() + string(os.PathSeparator) + "test.log"
	defer func() { logFile = "" }()

	initLog()
	var err error
	// test that input file is getting read
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
	inFile = "test.json"
	viper.Set("login", authString)
	os.Remove(inFile)
	var err error

	initLog()
	// test that absent input file throws an error
	tq.CaptureOutput(func() {
		err = test_cmd.Execute()
	})
	assert.Regexp(t, "cannot open input file .* for reading", err.Error())
	inFile = ""

}

// Test that execute returns errors and output and handles
func Test_Execute(t *testing.T) {
	rootCmd.SetArgs(nil)
	stdout, _ := tq.CaptureOutput(Execute)

	assert.Contains(t, string(stdout), "Usage:")

	initLog()
	_tq.SetOutput([]byte(`{"test":"json"}`))
	stdout, _ = tq.CaptureOutput(Execute)

	assert.Contains(t, ansi.Strip(string(stdout)), `{"test":"json"}`)

	pretty = true
	defer func() { pretty = false }()

	stdout, _ = tq.CaptureOutput(Execute)

	assert.Regexp(t, regexp.MustCompile(`\{\n\s+"test":\s+"json"\n\}`), ansi.Strip(string(stdout)))

}
