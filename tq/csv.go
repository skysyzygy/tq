package tq

import (
	"errors"

	"golang.org/x/text/transform"
)

// Adds additional quoting to the start and end of strings so that quotes are preserved
// after reading
type csvQuoter struct{}

func (c *csvQuoter) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	var b byte
	nQuotes := 0
	for nSrc < len(src) {
		if nDst+3 >= len(dst) {
			err = transform.ErrShortDst
			return
		}
		b = src[nSrc]
		// get number of quotes in a row
		if b == byte('"') {
			nQuotes++
			if nSrc+1 == len(src) && !atEOF {
				err = transform.ErrShortSrc
				return
			}
		}
		if b != byte('"') || nSrc+1 == len(src) && atEOF {
			if nQuotes%2 != 0 {
				dst[nDst] = '"'
				dst[nDst+1] = '"'
				nDst = nDst + 2
			}
			nQuotes = 0
		}
		dst[nDst] = b
		nDst++
		nSrc++
	}
	return
}
func (c *csvQuoter) Reset() {
}

// Convert a slice of jsonMap to a slice of csv records plus a header row ([]string)
func jsonMapsToCsv(in []jsonMap) (out csv) {
	keys := make(map[string]bool)
	out = make(csv, len(in)+1)

	// gather all keys
	for _, row := range in {
		for key := range row {
			keys[key] = true
		}
	}

	// fill in the csv
	i := 0
	out[0] = make([]string, len(keys))
	for key := range keys {
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
func jsonMapsFromCsv(in csv) (out []jsonMap, err error) {
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
