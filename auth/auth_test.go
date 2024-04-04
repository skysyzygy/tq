package auth

import (
	"os"
	"reflect"
	"testing"

	"bitbucket.org/TN_WebShare/gotess"
	"github.com/99designs/keyring"
	"github.com/stretchr/testify/assert"
)

// Setup the test environment by making a separate keystore for testing
func TestMain(m *testing.M) {
	// setup code
	keys, _ = keyring.Open(keyring.Config{
		ServiceName: "tq_test",
	})
	code := m.Run()
	// teardown code
	os.Exit(code)
}

func TestAuth_AuthString(t *testing.T) {
	assert := assert.New(t)

	v, err := Auth{"a", "b", "c", "d", nil}.AuthString()
	assert.Equal(v, "a:b:c:d", "generates string")
	assert.NoError(err)

	v, err = Auth{"a:", "b", "c", "d", nil}.AuthString()
	assert.Equal(v, "")
	assert.Error(err, ":", "complains when there are ':' present in hostname")

	v, err = Auth{"a", "b:", "c", "d", nil}.AuthString()
	assert.Equal(v, "")
	assert.Error(err, ":", "complains when there are ':' present in username")

	v, err = Auth{"a", "b", "c:", "d", nil}.AuthString()
	assert.Equal(v, "")
	assert.Error(err, ":", "complains when there are ':' present in usergroup")

	v, err = Auth{"a", "b", "c", "d:", nil}.AuthString()
	assert.Equal(v, "")
	assert.Error(err, ":", "complains when there are ':' present in location")

	v, err = Auth{"a", "b", "c", "d", []byte(":)")}.AuthString()
	assert.Equal(v, "a:b:c:d", "doesn't complain when there are ':' present in password")
	assert.NoError(err)

}

func TestAuthFromString(t *testing.T) {
	assert := assert.New(t)

	v, err := AuthFromString("a:b:c:d")
	assert.Equal(v, Auth{"a", "b", "c", "d", nil}, "parses string into Auth")
	assert.NoError(err)

	v, err = AuthFromString("a:b:c:d:e")
	assert.Equal(v, Auth{})
	assert.Error(err, "four", "complains when there are too many parts in the string")

	v, err = AuthFromString("a:b:c")
	assert.Equal(v, Auth{})
	assert.Error(err, "four", "complains when there are too few parts in the string")

}

func TestAuth_SaveAuth(t *testing.T) {
	assert := assert.New(t)

	err := Auth{"a", "b", "c", "d", []byte("e")}.SaveAuth()
	pass, _ := keys.Get("a:b:c:d")
	assert.Equal(pass.Data, []byte("e"), "saves Auth password to keystore")
	assert.NoError(err)

}

func TestAuth_LoadAuth(t *testing.T) {
	a := Auth{"a", "b", "c", "d", nil}

	a.LoadAuth()
	assert.Equal(t, a.password, []byte("e"), "loads Auth password from keystore")
}

func TestListAuths(t *testing.T) {
	v, err := ListAuths()

	assert.Equal(t, v, []Auth{{"a", "b", "c", "d", nil}}, "lists all auths in keystore")
	assert.NoError(t, err)

}

func TestAuth_DeleteAuth(t *testing.T) {
	k, _ := keys.Keys()
	assert.Equal(t, len(k), 1, "there's a key in the keystore")

	err := Auth{"a", "b", "c", "d", nil}.DeleteAuth()

	k, _ = keys.Keys()
	assert.Equal(t, len(k), 0, "deletes auth from keystore")
	assert.NoError(t, err)
}

func TestAuth_Login(t *testing.T) {
	tests := []struct {
		name    string
		a       Auth
		want    *gotess.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Login()
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
