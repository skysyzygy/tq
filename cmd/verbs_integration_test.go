package cmd

import (
	"testing"

	"github.com/99designs/keyring"
	"github.com/skysyzygy/tq/tq"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	viper.Set("login", "t-gw-test-b-ex-rest.bam.org/TessituraService/|jolson||jolson-14")
}

// end-to-end get test
func Test_Get_Integration(t *testing.T) {
	// use the integration keystore
	keys, _ = keyring.Open(keyring.Config{
		ServiceName: "tq_test_integration",
	})

	// test without payload
	rootCmd.SetArgs([]string{"get", "constituents"})
	err := rootCmd.Execute()
	assert.ErrorContains(t, err, "500")

	// test with invalid payload
	rootCmd.SetArgs([]string{"get", "constituents", `{"constituentId":"0"}`})
	err = rootCmd.Execute()
	assert.ErrorContains(t, err, "Constituent Id cannot be 0 or Null")

	// test with valid payload
	rootCmd.SetArgs([]string{"get", "constituents", `{"constituentId":"1"}`})
	stdout, stderr := tq.CaptureOutput(func() { err = rootCmd.Execute() })
	assert.NoError(t, err)
	assert.Regexp(t, "Using config file: .+tq\n$", string(stderr))
	assert.Contains(t, string(stdout), "Dummy")

}
