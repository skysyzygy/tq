package cmd

import (
	"testing"

	"github.com/99designs/keyring"
	"github.com/skysyzygy/tq/auth"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	viper.Set("login", "t-gw-test-b-ex-rest.bam.org/TessituraService|jolson||jolson-14")
}

// end-to-end get test
func Test_Get_Integration(t *testing.T) {
	// use the real keystore
	auth.Keys, _ = keyring.Open(keyring.Config{
		ServiceName: "tq",
	})

	// test without payload
	rootCmd.SetArgs([]string{"get", "constituents"})
	err := rootCmd.Execute()
	assert.ErrorContains(t, err, "sdf")

	// test with payload
	rootCmd.SetArgs([]string{"get", "constituents", `{"constituentId":"0"}`})
	err = rootCmd.Execute()
	assert.NoError(t, err)
}
