package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jsonHighlight(t *testing.T) {

	json := jsonHighlight(`{"Filter":"string","ID":"string","MaintenanceMode":"string"}`)
	// contains ANSI escape sequence
	assert.Contains(t, json, "\033")
}
