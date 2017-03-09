package dna

import (
	"fmt"
	"testing"
)

func TestRouterTreeSimple(t *testing.T) {
	r := newRootNode()
	r.addRoute("/user", []Handler{func(c *Context) {
		fmt.Println("/user")
	}})
	hs, _, _ := r.solve("/user")
	for _, h := range hs {
		h(&Context{})
	}
}

func TestRouterTreeParameter1(t *testing.T) {
	r := newRootNode()
	r.addRoute("/user/:name", []Handler{func(c *Context) {
		fmt.Println("/user/:name")
	}})
	hs, ps, _ := r.solve("/user/john")
	for _, h := range hs {
		h(&Context{})
	}
	fmt.Println(ps)

	hs, ps, _ = r.solve("/user/") // does not match !!
	for _, h := range hs {
		h(&Context{})
	}
	fmt.Println(ps)

	hs, ps, _ = r.solve("/user") // does not match !!
	for _, h := range hs {
		h(&Context{})
	}
	fmt.Println(ps)
}

func TestRouterTreeParameter2(t *testing.T) {
	r := newRootNode()
	r.addRoute("/user/*name", []Handler{func(c *Context) {
		fmt.Println("/user/*name")
	}})
	hs, ps, _ := r.solve("/user/john")
	for _, h := range hs {
		h(&Context{})
	}
	fmt.Println(ps)

	hs, ps, _ = r.solve("/user/") // match !!
	for _, h := range hs {
		h(&Context{})
	}
	fmt.Println(ps)

	hs, ps, _ = r.solve("/user") // does not match !!
	for _, h := range hs {
		h(&Context{})
	}
	fmt.Println(ps)
}

func TestRouterTreeComplex(t *testing.T) {
	r := newRootNode()
	r.addRoute("/user/*name/:action", []Handler{func(c *Context) {
		fmt.Println("/user/*name/:action")
	}})
	hs, ps, _ := r.solve("/user/john/send")
	for _, h := range hs {
		h(&Context{})
	}
	fmt.Println(ps)
}
