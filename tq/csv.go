package tq

import (
	"errors"
	"slices"

	"golang.org/x/exp/maps"
	"golang.org/x/text/transform"
)

// Adds additional quoting to the start and end of strings so that quotes are preserved
// after reading
type csvQuoter struct {
	nQuotes int
}

func (c *csvQuoter) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	var b byte
	for nSrc < len(src) {
		if nDst+3 >= len(dst) {
			err = transform.ErrShortDst
			return
		}
		b = src[nSrc]
		// get number of quotes in a row
		if b == byte('"') {
			c.nQuotes++
		}
		if b != byte('"') || nSrc+1 == len(src) {
			if !atEOF {
				err = transform.ErrShortSrc
				return
			}
			if c.nQuotes%2 != 0 {
				dst[nDst] = '"'
				dst[nDst+1] = '"'
				nDst = nDst + 2
			}
			c.Reset()
		}
		dst[nDst] = b
		nDst++
		nSrc++
	}
	return
}
func (c *csvQuoter) Reset() {
	c.nQuotes = 0
}

// Convert a slice of jsonMap to a slice of csv records plus a header row ([]string)
func jsonMapsToRecords(in []jsonMap) (out records) {
	keys := make(map[string]bool)
	out = make(records, len(in)+1)

	// gather all keys
	for _, row := range in {
		for key := range row {
			keys[key] = true
		}
	}

	keysSorted := maps.Keys(keys)
	slices.Sort(keysSorted)
	// fill in the csv
	i := 0
	out[0] = make([]string, len(keys))
	for _, key := range keysSorted {
		out[0][i] = key
		for j, row := range in {
			if out[j+1] == nil {
				out[j+1] = make([]string, len(keys))
			}
			out[j+1][i] = string(row[key])
		}
		i++
	}
	return
}

// Convert a slice of csv records plus a header row ([]string) to a slice of jsonMap
func jsonMapsFromRecords(in records) (out []jsonMap, err error) {
	if len(in) == 0 {
		return nil, errors.New("csv has no rows")
	}

	keys := in[0]
	out = make([]jsonMap, len(in)-1)

	for i := 1; i < len(in); i++ {
		out[i-1] = make(jsonMap)
		for j, key := range keys {
			out[i-1][key] = []byte(in[i][j])
		}
	}

	return
}
