package main

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that instantiateStructType recurses into a struct, navigating pointers, and initializes
// strings with "string" and everything else with zero type
func Test_instantiateStructType(t *testing.T) {
	type test struct {
		I int
		F float32
		B bool
		S string
	}
	type nest struct {
		P *test
		O test
	}
	pointer := new(nest)
	object := *pointer
	j, _ := json.Marshal(instantiateStructType(reflect.TypeOf(object)))
	assert.Equal(t,
		[]byte(`{"P":{"I":0,"F":0,"B":false,"S":"string"},"O":{"I":0,"F":0,"B":false,"S":"string"}}`),
		j)

	j, _ = json.Marshal(instantiateStructType(reflect.TypeOf(pointer)))
	assert.Equal(t,
		[]byte(`{"P":{"I":0,"F":0,"B":false,"S":"string"},"O":{"I":0,"F":0,"B":false,"S":"string"}}`),
		j)

}
