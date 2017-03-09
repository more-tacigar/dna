package dna

import (
	"net/http"
)

type Engine struct {
	Router *Router
}

func NewEngine() *Engine {
	return &Engine{
		Router: NewRouter(),
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Because dna is a framework for JSON API Server,
	// a response "Content-Type" must be "application/json".
	w.Header().Set("Content-Type", "application/json")

	m, err := Str2Method(r.Method)
	if err != nil {
		// If the method is invalid, responds 405 status code.
		w.WriteHeader(405)
		return
	}

	root, ok := e.Router.roots[m]
	if !ok {
		// If a router's root is not found, responds 404 status code.
		w.WriteHeader(404)
		return
	}

	handlers, urlParams, err := root.solve(r.URL.Path)
	if err != nil {
		// If not found path, responds 404 status code.
		w.WriteHeader(404)
		return
	}

	c := NewContext(w, r, urlParams)
	for ; c.index < len(handlers); c.index++ {
		handler := handlers[c.index]
		handler(c)
		if c.index >= abortIndex {
			break
		}
	}
}

func (e *Engine) Run(addr string) error {
	err := http.ListenAndServe(addr, e)
	return err
}
