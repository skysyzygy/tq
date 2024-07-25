package tq

import (
	"bytes"
	"encoding/json"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func jsonMapToStringMap(in jsonMap) (out map[string]string) {
	out = make(map[string]string)
	for key, value := range in {
		out[key] = string(value)
	}
	return
}

func Test_flattenJSONMap(t *testing.T) {
	j := jsonMap{
		"a": []byte(`"apple"`),
		"b": []byte(`[{"badger":"mammal"},{"banana":"fruit"},"bagel"]`),
		"c": []byte(`{"cucumber":"vegetable or fruit?"}`),
		"d": []byte(`null`),
		"e": []byte(`1`),
		"f": []byte(`false`),
	}
	f := jsonMap{
		"a":           []byte(`"apple"`),
		"b[0].badger": []byte(`"mammal"`),
		"b[1].banana": []byte(`"fruit"`),
		"b[2]":        []byte(`"bagel"`),
		"c.cucumber":  []byte(`"vegetable or fruit?"`),
		"d":           []byte(`null`),
		"e":           []byte(`1`),
		"f":           []byte(`false`),
	}
	jMarshaled, _ := json.Marshal(j)

	flattened, err := FlattenJSONMap(jMarshaled, "")
	assert.NoError(t, err)
	assert.Equal(t, jsonMapToStringMap(f), jsonMapToStringMap(flattened))
	assert.Equal(t, json.RawMessage(`1`), flattened["e"])
	assert.Equal(t, json.RawMessage(`false`), flattened["f"])
}

// Returns errors from each level
func Test_flattenJSONMapError(t *testing.T) {
	j := jsonMap{
		"a": []byte(`"apple"`),
		"b": []byte(`[{"badger":"mammal"},{"banana":"fruit"},"bagel"]`),
		"c": []byte(`{"cucumber":"vegetable or fruit?"}`),
		"d": []byte(`null`),
		"e": []byte(`1`),
		"f": []byte(`false`),
	}

	jMarshaled, _ := json.Marshal(j)

	jMarshaled = bytes.Replace(jMarshaled, j["b"], []byte(`[{"badger":"mammal"},{"banana":"fruit"}`), 1)
	_, err := FlattenJSONMap(jMarshaled, "")
	assert.ErrorContains(t, err, "invalid character")

	jMarshaled = bytes.Replace(jMarshaled, j["b"], []byte(`[{"badger":"mammal"},{"banana":"fruit"]`), 1)
	_, err = FlattenJSONMap(jMarshaled, "")
	assert.ErrorContains(t, err, "invalid character")

	jMarshaled = bytes.Replace(jMarshaled, j["b"], []byte(`[{"badger":"mammal"},{"banana":"fruit"}]`), 1)
	jMarshaled = bytes.Replace(jMarshaled, j["c"], []byte(`{"cucumber":"vegetable or fruit?"`), 1)
	_, err = FlattenJSONMap(jMarshaled, "")
	assert.ErrorContains(t, err, "invalid character")

}

// flattenJSONMap works when there is whitespace in the JSON
func Test_flattenJSONMapWhitespace(t *testing.T) {
	j := jsonMap{
		"a": []byte(` "apple" `),
		"b": []byte(` [ { "badger" : "mammal" } , { "banana" : "fruit" } ]`),
		"c": []byte(` { "cucumber" : "vegetable or fruit?" } `),
		"d": []byte(` null `),
		"e": []byte(` 1 `),
		"f": []byte(` false `),
	}
	f := jsonMap{
		"a":           []byte(`"apple"`),
		"b[0].badger": []byte(`"mammal"`),
		"b[1].banana": []byte(`"fruit"`),
		"c.cucumber":  []byte(`"vegetable or fruit?"`),
		"d":           []byte(`null`),
		"e":           []byte(`1`),
		"f":           []byte(`false`),
	}
	jMarshaled, _ := json.Marshal(j)
	flattened, err := FlattenJSONMap(jMarshaled, "")
	assert.NoError(t, err)
	assert.Equal(t, jsonMapToStringMap(f), jsonMapToStringMap(flattened))
	assert.Equal(t, json.RawMessage(`1`), flattened["e"])
	assert.Equal(t, json.RawMessage(`false`), flattened["f"])

}

func Test_unflattenJSONMap(t *testing.T) {
	j := jsonMap{
		"a": []byte(`"apple"`),
		"b": []byte(`[{"badger":"mammal"},null,{"banana":"fruit"},["bagel","beignet",[{"beagle":"dog"}]]]`),
		"c": []byte(`{"cucumber":"vegetable or fruit?"}`),
		"d": []byte(`null`),
		"e": []byte(`1`),
		"f": []byte(`false`),
	}
	f := jsonMap{
		"a":                 []byte(`"apple"`),
		"b[0].badger":       []byte(`"mammal"`),
		"b[2].banana":       []byte(`"fruit"`),
		"b[3][0]":           []byte(`"bagel"`),
		"b[3][1]":           []byte(`"beignet"`),
		"b[3][2][0].beagle": []byte(`"dog"`),
		"c.cucumber":        []byte(`"vegetable or fruit?"`),
		"d":                 []byte(`null`),
		"e":                 []byte(`1`),
		"f":                 []byte(`false`),
	}
	unflattened, err := unflattenJSONMap(f)
	assert.NoError(t, err)
	jMarshaled, _ := json.Marshal(j)
	assert.Equal(t, string(jMarshaled), string(unflattened))
}

func Test_unflattenJSONMapError(t *testing.T) {
	f := jsonMap{
		"a":           []byte(`"apple"`),
		"b[0].badger": []byte(`"mammal"`),
		"b[1].banana": []byte(`"fruit"`),
		"c.cucumber":  []byte(`"vegetable or fruit?"`),
		"d":           []byte(`null`),
		"e":           []byte(`1`),
		"f":           []byte(`false`),
	}

	_, err := unflattenJSONMap(f)
	assert.NoError(t, err)

	f["b[A].bah"] = []byte(`"sheep"`)
	_, err = unflattenJSONMap(f)
	assert.Error(t, err)
	assert.Regexp(t, "(?s)key b.+invalid syntax", err.Error())

	delete(f, "b[A].bah")
	f["j[5].jaberwocky"] = []byte(`{"doesn't exist":`)
	_, err = unflattenJSONMap(f)
	assert.Error(t, err)
	assert.Regexp(t, "(?s)key j.+5.+unexpected end of JSON", err.Error())

}

func Test_updateJSONMap(t *testing.T) {
	a := jsonMap{
		"one":   []byte("two"),
		"three": []byte("four"),
	}
	b := jsonMap{
		"one": []byte("one"),
		"two": []byte("two"),
	}

	updateJSONMap(a, b)

	assert.Equal(t, "two", string(b["one"]))
	assert.Equal(t, "two", string(b["two"]))
	assert.Equal(t, "four", string(b["three"]))
}

func Test_jsonMapsToCsv(t *testing.T) {
	a := []jsonMap{{
		`"one key with spaces"`:            []byte(`"one"`),
		"another[one].with.control[chars]": []byte(`2`),
		"nil":                              []byte(`null`),
		"not":                              []byte(`false`),
	}}
	as := slices.Concat(a, a, a, a)

	csv := jsonMapsToRecords(as)

	assert.Equal(t, 5, len(csv))
	assert.Equal(t, 4, len(csv[0]))
	assert.Equal(t, []string{"\"one key with spaces\"", "another[one].with.control[chars]", "nil", "not"}, csv[0])
	assert.Equal(t, []string{"\"one\"", "2", "null", "false"}, csv[1])
}

func Test_jsonMapsFromCsv(t *testing.T) {
	a := records{{
		"one key with spaces",
		"another[one].with.control[chars]",
		"nil",
		"not",
	}, {`"one"`,
		`2`,
		`null`,
		`false`,
	}}
	b := []jsonMap{{
		"another[one].with.control[chars]": []byte(`2`),
		"nil":                              []byte(`null`),
		"not":                              []byte(`false`),
		"one key with spaces":              []byte(`"one"`),
	}}
	as := slices.Concat(a, a[1:], a[1:])
	bs := slices.Concat(b, b, b)

	js, _ := jsonMapsFromRecords(as)

	assert.Equal(t, len(as)-1, len(js))
	assert.Equal(t, bs, js)
}
