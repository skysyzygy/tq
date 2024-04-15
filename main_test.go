package main

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/skysyzygy/tq/auth"
	"github.com/skysyzygy/tq/client/g_e_t"
	"github.com/skysyzygy/tq/models"
	"github.com/stretchr/testify/assert"
)

func Test_unmarshallStructWithRemainder(t *testing.T) {
	type P struct{ A, B, C string }
	type Q map[string]any

	// test that unmarshallWithRemainder fills struct and returns map of extra data
	p := new(P)
	res, err := unmarshallStructWithRemainder([]byte(`{"A": "these", "B": "are", "C": "words"}`), p)
	assert.Equal(t, P{"these", "are", "words"}, *p)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))

	p = new(P)
	res, err = unmarshallStructWithRemainder([]byte(`{"A": "these", "B": "are", "C": "words", "D": "nothing", "E": "more"}`), p)
	assert.Equal(t, P{"these", "are", "words"}, *p)
	assert.NoError(t, err)
	assert.Equal(t, Q{"D": "nothing", "E": "more"}, Q(res))

	// test that unmarshallWithRemainder returns an error if it's asked to fill a non-struct
	q := new(Q)
	res, err = unmarshallStructWithRemainder([]byte(`{"A": "these", "B": "are", "C": "words"}`), q)
	assert.Equal(t, Q(nil), *q)
	assert.ErrorContains(t, err, "must be pointer to struct")
	assert.Equal(t, Q(nil), Q(res))

	// test that unmarshallWithRemainder returns an error if it's given a JSON array
	p = new(P)
	res, err = unmarshallStructWithRemainder([]byte(`[{"A": "these", "B": "are", "C": "words"}]`), p)
	assert.ErrorContains(t, err, "cannot unmarshal array")
	assert.Equal(t, Q(nil), Q(res))

}

func Test_unmarshallNestedStructWithRemainder(t *testing.T) {
	type P struct{ A, B, C string }
	type N struct {
		D, E, F string
		Nest    P
	}

	// test that unmarshallNestedWithRemainder fills nested struct and returns map of extra data
	n := new(N)
	res, err := unmarshallNestedStructWithRemainder([]byte(`{"A": "these", "B": "are", "C": "words"}`), n)
	assert.Equal(t, N{Nest: P{"these", "are", "words"}}, *n)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))

	n = new(N)
	res, err = unmarshallNestedStructWithRemainder([]byte(`{"A": "these", "B": "are", "C": "words", "D": "nothing", "E": "more"}`), n)
	assert.Equal(t, N{D: "nothing", E: "more", Nest: P{"these", "are", "words"}}, *n)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))

	n = new(N)
	res, err = unmarshallNestedStructWithRemainder([]byte(`{"A": "these", "B": "are", "C": "words", "D": "nothing", "Z": "zzz"}`), n)
	assert.Equal(t, N{D: "nothing", Nest: P{"these", "are", "words"}}, *n)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, map[string]any{"Z": "zzz"}, res)

}

// test that structFields returns used elements of struct pointer, string, and int types
func Test_structFields(t *testing.T) {

	strings := []string{"these", "are", "words"}
	pp := struct{ A, B, C, empty *string }{&strings[0], &strings[1], &strings[2], nil}
	ps := struct{ A, B, C, empty string }{strings[0], strings[1], strings[2], ""}
	pi := struct{ A, B, C, empty int }{1, 2, 3, 0}

	assert.Equal(t, []string{"A", "B", "C"}, structFields(pp))
	assert.Equal(t, []string{"A", "B", "C"}, structFields(ps))
	assert.Equal(t, []string{"A", "B", "C"}, structFields(pi))

}

// test that Login builds Tessitura API client and saves BasicAuth info for future use
func Test_Login(t *testing.T) {
	tq := new(tqConfig)
	tq.Login(auth.New("hostname", "user", "", "", nil))
	assert.NotNil(t, tq.Get)
	assert.NotNil(t, tq.basicAuth)
}

// test that DoOne calls swagger API functions and returns a response
func Test_DoOne(t *testing.T) {
	oneConstituent := models.Constituent{
		ID:        0,
		FirstName: "Test",
		LastName:  "User",
	}
	oneConstituentDetail := models.ConstituentDetail{
		ID:        0,
		FirstName: "Test",
		LastName:  "User",
	}
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := g_e_t.ConstituentsGetConstituentParams{}

		// Check that the caller is authenticated
		assert.Equal(t, "Basic "+base64.StdEncoding.EncodeToString([]byte(`user:::password`)),
			r.Header.Values("Authorization")[0])

		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, &req)

		resBody, _ := json.Marshal(oneConstituent)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resBody)
	}))
	defer server.Close()
	tq := new(tqConfig)
	query := []byte(`{"Id": 0}`)
	tq.Login(auth.New(strings.Replace(server.URL, "https://", "", 1), "user", "", "", []byte("password")))

	res, err := DoOne(*tq, tq.Get.ConstituentsGet, query)
	expectedJSON, _ := json.Marshal(oneConstituent)
	assert.Equal(t, expectedJSON, res)
	assert.NoError(t, err)

	res, err = DoOne(*tq, tq.Put.ConstituentsUpdate, query)
	expectedJSON, _ = json.Marshal(oneConstituent)
	assert.Equal(t, expectedJSON, res)
	assert.NoError(t, err)

	res, err = DoOne(*tq, tq.Post.ConstituentsCreateConstituent, query)
	expectedJSON, _ = json.Marshal(oneConstituentDetail)
	assert.Equal(t, expectedJSON, res)
	assert.NoError(t, err)

}
