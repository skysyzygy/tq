package main

import (
	"io"
	"os"
	"regexp"
	"testing"

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
	assert.ErrorContains(t, err, "must be struct")
	assert.Equal(t, Q(nil), Q(res))

	// test that unmarshallWithRemainder returns an error if it's given a JSON array
	p = new(P)
	res, err = unmarshallStructWithRemainder([]byte(`[{"A": "these", "B": "are", "C": "words"}]`), p)
	assert.ErrorContains(t, err, "cannot unmarshal array")
	assert.Equal(t, Q(nil), Q(res))

}

// test that unmarshalLStructReport returns used and unused element of struct and map for pointer, string, and int types
func Test_unmarshallStructReport(t *testing.T) {

	strings := []string{"these", "are", "words"}
	pp := struct{ A, B, C, empty *string }{&strings[0], &strings[1], &strings[2], nil}
	q := map[string]any{"D": "nothing", "E": "more"}

	r, w, _ := os.Pipe()
	unmarshallStructReport(pp, q, w)
	output, err := io.ReadAll(r)
	assert.Regexp(t, regexp.MustCompile(`executing a query using fields \[[ABC ]+\] and ignoring fields \[[DE ]+\]`), string(output))
	assert.NoError(t, err)

	ps := struct{ A, B, C, empty string }{strings[0], strings[1], strings[2], ""}
	r, w, _ = os.Pipe()
	unmarshallStructReport(ps, q, w)
	output, err = io.ReadAll(r)
	assert.Regexp(t, regexp.MustCompile(`executing a query using fields \[[ABC ]+\] and ignoring fields \[[DE ]+\]`), string(output))
	assert.NoError(t, err)

	pi := struct{ A, B, C, empty int }{1, 2, 3, 0}
	r, w, _ = os.Pipe()
	unmarshallStructReport(pi, q, w)
	output, err = io.ReadAll(r)
	assert.Regexp(t, regexp.MustCompile(`executing a query using fields \[[ABC ]+\] and ignoring fields \[[DE ]+\]`), string(output))
	assert.NoError(t, err)
}
