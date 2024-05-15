package cmd

import (
	"os"
	"testing"

	"github.com/skysyzygy/tq/auth"

	"github.com/skysyzygy/tq/tq"
	"github.com/stretchr/testify/assert"
)

var skipAuth bool = false

// authenticateAddCmd adds an auth method
func Test_authenticateAddCmd(t *testing.T) {
	if skipAuth {
		t.Skip()
	}
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

	key, err := auth.Keys.Get("tessitura.api/basePath|user|group|location")
	assert.Error(t, err)

	authenticateAddCmd.Run(authenticateAddCmd, nil)

	key, err = auth.Keys.Get("tessitura.api/basePath|user|group|location")
	assert.Equal(t, "tessitura.api/basePath|user|group|location", key.Key)
	assert.Equal(t, []byte("password"), key.Data)
	assert.NoError(t, err)
}

// authenticateListCmd lists all authentication methods
func Test_authenticateListCmd(t *testing.T) {
	if skipAuth {
		t.Skip()
	}

	*hostname = ""
	*username = ""
	*usergroup = ""
	*location = ""

	stdout, stderr := tq.CaptureOutput(func() {
		authenticateListCmd.Run(authenticateListCmd, nil)
	})

	assert.Contains(t, string(stdout), "{tessitura.api/basePath user group location }")
	assert.Equal(t, []byte{}, stderr)
}

// authenticateDeleteCmd removes an authentication method
func Test_authenticateDeleteCmd(t *testing.T) {
	if skipAuth {
		t.Skip()
	}

	*hostname = "tessitura.api/basePath"
	*username = "user"
	*usergroup = "group"
	*location = "location"

	_, err := auth.Keys.Get("tessitura.api/basePath|user|group|location")
	assert.NoError(t, err)

	authenticateDeleteCmd.Run(authenticateDeleteCmd, nil)

	_, err = auth.Keys.Get("tessitura.api/basePath|user|group|location")
	assert.Error(t, err)
}
