package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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
	assert.Equal(v, "a|b|c|d", "generates string")
	assert.NoError(err)

	v, err = Auth{"a|", "b", "c", "d", nil}.AuthString()
	assert.Equal(v, "")
	assert.ErrorContains(err, "|", "complains when there are '|' present in hostname")

	v, err = Auth{"a", "b|", "c", "d", nil}.AuthString()
	assert.Equal(v, "")
	assert.ErrorContains(err, "|", "complains when there are '|' present in username")

	v, err = Auth{"a", "b", "c|", "d", nil}.AuthString()
	assert.Equal(v, "")
	assert.ErrorContains(err, "|", "complains when there are '|' present in usergroup")

	v, err = Auth{"a", "b", "c", "d|", nil}.AuthString()
	assert.Equal(v, "")
	assert.ErrorContains(err, "|", "complains when there are '|' present in location")

	v, err = Auth{"a", "b", "c", "d", []byte("|)")}.AuthString()
	assert.Equal(v, "a|b|c|d", "doesn't complain when there are '|' present in password")
	assert.NoError(err)

}

func TestAuthFromString(t *testing.T) {
	assert := assert.New(t)

	v, err := AuthFromString("a|b|c|d")
	assert.Equal(v, Auth{"a", "b", "c", "d", nil}, "parses string into Auth")
	assert.NoError(err)

	v, err = AuthFromString("a|b|c|d|e")
	assert.Equal(v, Auth{})
	assert.Error(err, "four", "complains when there are too many parts in the string")

	v, err = AuthFromString("a|b|c")
	assert.Equal(v, Auth{})
	assert.Error(err, "four", "complains when there are too few parts in the string")

}

func TestAuth_SaveAuth(t *testing.T) {
	assert := assert.New(t)

	err := Auth{"a", "b", "c", "d", []byte("e")}.SaveAuth()
	pass, _ := keys.Get("a|b|c|d")
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

func TestAuth_Validate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := gotess.AuthenticationRequest{}
		var (
			isAuthenticated bool
			message         string
		)

		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, &req)

		if r.RequestURI != "/Security/Authenticate/" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if *req.UserName == "user" && *req.UserGroup == "group" &&
			*req.Password == "password" && *req.MachineLocation == "location" {
			isAuthenticated = true
		} else {
			isAuthenticated = false
			message = "bad credentials"
		}

		res := gotess.AuthenticationResponse{}
		res.IsAuthenticated = &isAuthenticated
		res.Message = &message

		resBody, _ := json.Marshal(res)
		w.Write(resBody)
	}))
	defer server.Close()

	v, err := Auth{hostname: "https://not-a-host.com",
		username: "user", usergroup: "group", location: "location", password: []byte("password")}.Validate()
	assert.False(t, v)
	assert.ErrorContains(t, err, "no such host", "validation fails when server is unreachable and provides useful info")

	v, err = Auth{hostname: server.URL + "/Not an endpoint/",
		username: "user", usergroup: "group", location: "location", password: []byte("password")}.Validate()
	assert.False(t, v)
	assert.ErrorContains(t, err, "404 Not Found", "validation fails when endpoint is incorrect and provides useful info")

	v, err = Auth{hostname: server.URL,
		username: "user", usergroup: "group", location: "location", password: []byte("password")}.Validate()
	assert.True(t, v, "validation works when credentials are correct")
	assert.NoError(t, err)

	v, err = Auth{hostname: server.URL,
		username: "user", usergroup: "group", location: "location", password: []byte("wrongPass")}.Validate()
	assert.False(t, v)
	assert.ErrorContains(t, err, "bad credentials", "validation failes when credentials are incorrect")

}

func TestAuth_Validate_Integration(t *testing.T) {

	keys, _ = keyring.Open(keyring.Config{
		ServiceName: "tq_test_integration",
	})

	auths, _ := ListAuths()
	a := auths[0]
	a.LoadAuth()

	a1 := a
	a1.hostname = "https://not-a-server.bam.org/"
	v, err := a1.Validate()
	assert.False(t, v)
	assert.ErrorContains(t, err, "no such host", "validation fails when server is unreachable and provides useful info")

	a2 := a
	a2.hostname = strings.ReplaceAll(a2.hostname, "/TessituraService", "")
	v, err = a2.Validate()
	assert.False(t, v)
	assert.ErrorContains(t, err, "404 Not Found", "validation fails when endpoint is incorrect and provides useful info")

	v, err = a.Validate()
	assert.True(t, v, "validation works when credentials are correct")
	assert.NoError(t, err)

	a.password = []byte("wrong_password")
	v, err = a.Validate()
	assert.False(t, v)
	assert.ErrorContains(t, err, "Invalid Credentials", "validation fails when credentials are incorrect")

}

// func TestAuth_Login(t *testing.M)
