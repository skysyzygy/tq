package tq

import (
	csvReader "encoding/csv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/transform"
)

func Test_csvQuoter(t *testing.T) {
	in := strings.NewReader(`"this needs to be quoted",this,doesn't,"this has ""quotes""","so does this"`)

	q := transform.NewReader(in, &csvQuoter{})
	c := csvReader.NewReader(q)

	// qs, qerr := io.ReadAll(q)
	// fmt.Println(string(qs), qerr)
	cs, cerr := c.ReadAll()
	assert.Equal(t, []string{`"this needs to be quoted"`, "this", "doesn't", `"this has "quotes""`, `"so does this"`}, cs[0])
	assert.NoError(t, cerr)

}
