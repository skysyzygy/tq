package auth

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/99designs/keyring"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/skysyzygy/tq/client"
	"github.com/skysyzygy/tq/client/p_o_s_t"
	"github.com/skysyzygy/tq/models"
)

// Simple structure to hold authentication information
type Auth struct {
	// hostname and basepath without the leading https://
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
	if strings.Contains(a.hostname, "|") ||
		strings.Contains(a.username, "|") ||
		strings.Contains(a.usergroup, "|") ||
		strings.Contains(a.location, "|") {
		return "", errors.New("authentication info can't contain the '|' character")
	}
	return fmt.Sprintf("%v|%v|%v|%v", a.hostname, a.username, a.usergroup, a.location), nil
}

// Build the authentication string used for basic http authentication
func (a Auth) BasicAuth() (func(*runtime.ClientOperation), error) {
	if strings.Contains(a.username, ":") ||
		strings.Contains(a.usergroup, ":") ||
		strings.Contains(a.location, ":") {
		return nil, errors.New("authentication info can't contain the ':' character for Basic auth")
	}

	return func(co *runtime.ClientOperation) {
		co.AuthInfo = httptransport.BasicAuth(fmt.Sprintf("%v:%v:%v", a.username, a.usergroup, a.location), string(a.password))
	}, nil
}

// Parse an authentication string (i.e. from AuthString) into an Auth struct
func AuthFromString(str string) (Auth, error) {
	strs := strings.Split(str, "|")
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

// Validate authentication with the Tessitura API server at a.hostname
func (a Auth) Validate() (bool, error) {
	host := append(strings.SplitN(a.hostname, "/", 2), "")
	ignoreCerts, _ := httptransport.TLSClient(httptransport.TLSClientOptions{
		InsecureSkipVerify: true,
	})
	transport := httptransport.NewWithClient(host[0], host[1], []string{"https"}, ignoreCerts)
	client := client.New(transport, nil)

	request := p_o_s_t.AuthenticateAuthenticateParams{
		AuthenticationRequest: &models.AuthenticationRequest{
			UserName:        a.username,
			UserGroup:       a.usergroup,
			MachineLocation: a.location,
			Password:        string(a.password),
		}}

	response, err := client.Post.AuthenticateAuthenticate(&request)

	if err != nil {
		var dnsError *net.DNSError
		var apiError *runtime.APIError
		if errors.As(err, &dnsError) {
			return false, dnsError
		} else if errors.As(err, &apiError) {
			return false, fmt.Errorf("login failed with http response: %v",
				apiError.Response.(runtime.ClientResponse).Message(),
			)
		} else {
			return false, err
		}
	} else if response.IsSuccess() && response.Payload != nil {
		if response.Payload.IsAuthenticated {
			// Successful login!
			return true, nil
		} else {
			return false, fmt.Errorf("login failed with Tessitura response: %v", response.Payload.Message)
		}
	}
	// Should never get here but who knows?
	return false, fmt.Errorf("login failed with unknown error: %v %v",
		response.Code(), response.Error())
}
