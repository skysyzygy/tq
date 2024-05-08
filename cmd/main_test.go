package cmd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/keyring"

	"github.com/skysyzygy/tq/auth"
	"github.com/skysyzygy/tq/models"

	"encoding/json"
)

var authString string

func TestMain(m *testing.M) {
	server := server()
	defer server.Close()

	// Setup the test environment by make a separate keystore for testing
	auth.Keys, _ = keyring.Open(keyring.Config{
		ServiceName: "tq_test",
	})

	a := auth.New(strings.Replace(server.URL, "https://", "", 1), "", "", "", nil)
	a.Save()
	authString, _ = a.String()

	m.Run()

	a.Delete()
}

func server() (server *httptest.Server) {
	server = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resBody []byte
		w.Header().Set("Content-Type", "application/json")

		if r.RequestURI == "/Security/Authenticate" {
			res := models.AuthenticationResponse{
				IsAuthenticated: true,
				Message:         "",
			}
			resBody, _ = json.Marshal(&res)
		}

		w.Write(resBody)
	}))
	return
}
