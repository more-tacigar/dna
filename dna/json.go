package dna

import (
	"io"

	"github.com/tacigar/dna/json"
)

func DecodeJson(v interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(v)
}

func EncodeJson(v interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(v)
}
