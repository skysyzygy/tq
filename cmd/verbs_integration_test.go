package cmd

import (
	"os"
	"testing"

	"github.com/99designs/keyring"
	"github.com/skysyzygy/tq/tq"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	viper.Set("login", "t-gw-test-b-ex-rest.bam.org/TessituraService/|jolson||jolson-14")
}

func sendToStdin(query string) {
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(query))
	w.Close()
}

// end-to-end get test
func Test_Get_Integration_empty(t *testing.T) {
	// use the integration keystore
	keys, _ = keyring.Open(keyring.Config{
		ServiceName: "tq_test_integration",
	})

	// test without payload
	rootCmd.SetArgs([]string{"get", "constituents"})
	err := rootCmd.Execute()
	assert.ErrorContains(t, err, "500")
}

func Test_Get_Integration_invalid(t *testing.T) {
	// test with invalid payload
	rootCmd.SetArgs([]string{"get", "constituents"})
	sendToStdin(`{"constituentId":"0"}`)
	err := rootCmd.Execute()
	assert.ErrorContains(t, err, "Constituent Id cannot be 0 or Null")
}

func Test_Get_Integration_valid(t *testing.T) {
	var err error
	// test with valid payload
	rootCmd.SetArgs([]string{"get", "constituents"})
	sendToStdin(`{"constituentId":"1"}`)
	stdout, stderr := tq.CaptureOutput(func() { err = rootCmd.Execute() })
	assert.NoError(t, err)
	assert.Regexp(t, "Using config file: .+tq\n$", string(stderr))
	assert.Contains(t, string(stdout), "Dummy")
}
