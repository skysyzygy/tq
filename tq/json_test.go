package tq

import (
	"encoding/json"
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
		"b": []byte(`[{"badger":"mammal"},{"banana":"fruit"}]`),
		"c": []byte(`{"cucumber":"vegetable or fruit?"}`),
		"d": []byte(`null`),
		"e": []byte(`1`),
		"f": []byte(`false`),
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
	flattened, err := flattenJSONMap(j, "")
	assert.NoError(t, err)
	assert.Equal(t, jsonMapToStringMap(f), jsonMapToStringMap(flattened))
	assert.Equal(t, json.RawMessage(`1`), flattened["e"])
	assert.Equal(t, json.RawMessage(`false`), flattened["f"])
}

// Returns errors from each level
func Test_flattenJSONMapError(t *testing.T) {
	j := jsonMap{
		"a": []byte(`"apple"`),
		"b": []byte(`[{"badger":"mammal"},{"banana":"fruit"}]`),
		"c": []byte(`{"cucumber":"vegetable or fruit?"}`),
		"d": []byte(`null`),
		"e": []byte(`1`),
		"f": []byte(`false`),
	}
	_, err := flattenJSONMap(j, "")
	assert.NoError(t, err)

	j["b"] = []byte(`[{"badger":"mammal"},{"banana":"fruit"}`)
	_, err = flattenJSONMap(j, "")
	assert.ErrorContains(t, err, "unexpected end")

	j["b"] = []byte(`[{"badger":"mammal"},{"banana":"fruit"]`)
	_, err = flattenJSONMap(j, "")
	assert.ErrorContains(t, err, "]")

	j["b"] = []byte(`[{"badger":"mammal"},{"banana":"fruit"}]`)
	j["c"] = []byte(`{"cucumber":"vegetable or fruit?"`)
	_, err = flattenJSONMap(j, "")
	assert.ErrorContains(t, err, "unexpected end")

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
	flattened, err := flattenJSONMap(j, "")
	assert.NoError(t, err)
	assert.Equal(t, jsonMapToStringMap(f), jsonMapToStringMap(flattened))
	assert.Equal(t, json.RawMessage(`1`), flattened["e"])
	assert.Equal(t, json.RawMessage(`false`), flattened["f"])

}

func Test_unflattenJSONMap(t *testing.T) {
	j := jsonMap{
		"a": []byte(`"apple"`),
		"b": []byte(`[{"badger":"mammal"},{"banana":"fruit"}]`),
		"c": []byte(`{"cucumber":"vegetable or fruit?"}`),
		"d": []byte(`null`),
		"e": []byte(`1`),
		"f": []byte(`false`),
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
	unflattened, err := unflattenJSONMap(f)
	assert.NoError(t, err)
	assert.Equal(t, jsonMapToStringMap(j), jsonMapToStringMap(unflattened))
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

	f["b[a].bah"] = []byte(`"sheep"`)
	_, err = unflattenJSONMap(f)
	assert.Regexp(t, "(?s)key b\\[a\\].+invalid syntax", err.Error())

	delete(f, "b[a]")
	f["j[5].jaberwocky"] = []byte(`"doesn't exist"`)
	_, err = unflattenJSONMap(f)
	assert.Regexp(t, "(?s)key j\\[5\\].+index out of range", err.Error())

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
