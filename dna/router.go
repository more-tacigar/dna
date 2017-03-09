package dna

type Router struct {
	roots  map[Method]*node
	before []Handler
	after  []Handler
}

func NewRouter() *Router {
	return &Router{
		roots:  make(map[Method]*node),
		before: []Handler{},
		after:  []Handler{},
	}
}

func (r *Router) AddHandler(method Method, path string, handler Handler) error {
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = newRootNode()
	}
	hs := []Handler{}
	hs = append(hs, r.before...)
	hs = append(hs, handler)
	hs = append(hs, r.after...)
	return r.roots[method].addRoute(path, hs)
}

func (r *Router) GET(path string, handler Handler) error {
	return r.AddHandler(GET, path, handler)
}

func (r *Router) POST(path string, handler Handler) error {
	return r.AddHandler(POST, path, handler)
}

func (r *Router) DELETE(path string, handler Handler) error {
	return r.AddHandler(DELETE, path, handler)
}

func (r *Router) UPDATE(path string, handler Handler) error {
	return r.AddHandler(UPDATE, path, handler)
}

func (r *Router) Before(handler Handler) {
	r.before = append(r.before, handler)
}

func (r *Router) After(handler Handler) {
	r.after = append(r.after, handler)
}

func (r *Router) Branch() *Router {
	rr := NewRouter()
	copy(rr.before, r.before)
	copy(rr.after, r.after)
	return rr
}
