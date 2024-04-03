package auth

import (
	"errors"
	"fmt"
	"strings"

	"bitbucket.org/TN_WebShare/gotess"
	"github.com/99designs/keyring"
)

// Simple structure to hold authentication information
type Auth struct {
	hostname  string
	username  string
	usergroup string
	location  string
	password  []byte
}

var (
	keys keyring.Keyring
)

func init() {
	var err error
	keys, err = keyring.Open(keyring.Config{
		ServiceName: "tq",
	})
	if err != nil {
		panic(err)
	}
}

// Build the authentication string used for storing the password in the keystore
func (a Auth) AuthString() (string, error) {
	if strings.Contains(a.hostname, ":") ||
		strings.Contains(a.username, ":") ||
		strings.Contains(a.usergroup, ":") ||
		strings.Contains(a.location, ":") {
		return "", errors.New("authentication info can't contain the ':' character")
	}
	return fmt.Sprintf("%v:%v:%v:%v", a.hostname, a.username, a.usergroup, a.location), nil
}

// Parse an authentication string (i.e. from AuthString) into an Auth struct
func AuthFromString(str string) (Auth, error) {
	strs := strings.Split(str, ":")
	if len(strs) != 4 {
		return Auth{}, errors.New("authentication string must have exactly four values")
	}
	return Auth{strs[0], strs[1], strs[2], strs[3], nil}, nil
}

// Save the authentication data in the keystore
func (a Auth) SaveAuth() error {
	if authString, err := a.AuthString(); err != nil {
		return err
	} else {
		return keys.Set(keyring.Item{Key: authString, Data: a.password})
	}
}

// Load the password for the matching authentication in the keystore
func (a *Auth) LoadAuth() error {
	if authString, err := a.AuthString(); err != nil {
		return err
	} else {
		pass, err := keys.Get(authString)
		a.password = pass.Data
		return err
	}
}

// Delete the matching authentication in the keystore
func (a Auth) DeleteAuth() error {
	if authString, err := a.AuthString(); err != nil {
		return err
	} else {
		return keys.Remove(authString)
	}
}

// List all authentication keys in the keystore
func ListAuths() ([]Auth, error) {
	keys, err := keys.Keys()
	auths := make([]Auth, len(keys))
	for i, key := range keys {
		var err2 error
		auths[i], err2 = AuthFromString(key)
		err = errors.Join(err, err2)
	}
	return auths, err
}

// Log in to Tessitura with the given authentication and return a pointer to a client struct for use
func (a Auth) Login() (*gotess.Client, error) {
	client := gotess.NewClient(nil)
	client.SetBaseURL(a.hostname)
	client.SetCredentials(a.username, a.usergroup, a.location, string(a.password))

	status, err := client.Diagnostics.GetStatus()
	if err == nil && !*status.Success {
		err = errors.New("login failed")
	}
	return client, err
}
