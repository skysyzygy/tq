package cmd

import (
	"os"
	"testing"

	"github.com/skysyzygy/tq/auth"

	"github.com/99designs/keyring"
	"github.com/skysyzygy/tq/tq"
	"github.com/stretchr/testify/assert"
)

// Setup the test environment by making a separate keystore for testing
func TestMain(m *testing.M) {
	// setup code
	auth.Keys, _ = keyring.Open(keyring.Config{
		ServiceName: "tq_test",
	})
	code := m.Run()
	// teardown code
	os.Exit(code)
}

// authenticateAddCmd adds an auth method
func Test_authenticateAddCmd(t *testing.T) {
	// test that the https gets removed
	*hostname = "https://tessitura.api/basePath"
	*username = "user"
	*usergroup = "group"
	*location = "location"
	password := []byte("password")

	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write(password)
	w.Close()

	authenticateAddCmd.Run(authenticateAddCmd, nil)

	keys, _ := keyring.Open(keyring.Config{
		ServiceName: "tq",
	})
	ks, _ := keys.Keys()
	assert.Len(t, ks, 1)
	assert.Equal(t, "tessitura.api/basePath|user|group|location", ks[0])

	pass, _ := keys.Get(ks[0])
	assert.Equal(t, []byte("password"), pass.Data)
}

// authenticateListCmd lists all authentication methods
func Test_authenticateListCmd(t *testing.T) {
	*hostname = ""
	*username = ""
	*usergroup = ""
	*location = ""

	stdout, stderr := tq.CaptureOutput(func() {
		authenticateListCmd.Run(authenticateListCmd, nil)
	})

	assert.Equal(t, "{tessitura.api/basePath user group location }", string(stdout))
	assert.Equal(t, []byte{}, stderr)
}

// authenticateDeleteCmd removes an authentication method
func Test_authenticateDeleteCmd(t *testing.T) {
	*hostname = "tessitura.api/basePath"
	*username = "user"
	*usergroup = "group"
	*location = "location"

	authenticateDeleteCmd.Run(authenticateDeleteCmd, nil)

	keys, _ := keyring.Open(keyring.Config{
		ServiceName: "tq",
	})
	ks, _ := keys.Keys()
	assert.Len(t, ks, 0)
}
