package dna

import (
	"bytes"
	"testing"
)

func TestBefore(t *testing.T) {
	r := NewRouter()
	b := new(bytes.Buffer)
	r.Before(func(c *Context) {
		b.WriteRune('a')
	})
	r.Before(func(c *Context) {
		b.WriteRune('b')
	})
	r.After(func(c *Context) {
		b.WriteRune('d')
	})
	r.GET("/api", func(c *Context) {
		b.WriteRune('c')
	})
	hs, _, err := r.roots[GET].solve("/api")
	if err != nil {
		t.Error(err)
	}
	for _, h := range hs {
		h(&Context{})
	}
	if b.String() != "abcd" {
		t.Errorf("b expected abcd but %s", b.String())
	}
}
