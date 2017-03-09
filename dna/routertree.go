package dna

import (
	"strings"
)

type node struct {
	name     string
	children []*node
	handlers []Handler
}

func newNode(name string) *node {
	return &node{
		name:     name,
		children: []*node{},
		handlers: []Handler{},
	}
}

func newRootNode() *node {
	return newNode("")
}

func (n *node) findChild(name string) (*node, bool) {
	for _, c := range n.children {
		if c.name == name {
			return c, true
		}
	}
	return nil, false
}

func (n *node) addRoute(relativePath string, handlers []Handler) error {
	currnode := n
	segs := strings.Split(relativePath, "/")

	// If the size of segs is less than 2, the given relative path does not start
	// "/", so the path is invalid.
	if len(segs) <= 1 {
		return ErrInvalidPath
	}

	// If the size of segs is 2 and segs[1] is empty string, the given relative
	// path is "/", so register handler to `n` and return this function.
	if len(segs) == 2 && segs[1] == "" {
		currnode.handlers = append(currnode.handlers, handlers...)
		return nil
	}
	for i := 1; i < len(segs); i++ {
		seg := segs[i]
		c, ok := currnode.findChild(seg)
		if ok {
			currnode = c
		} else {
			// If not find a new node, move currnode to there.
			newnode := newNode(seg)
			currnode.children = append(currnode.children, newnode)
			currnode = newnode
		}
	}
	currnode.handlers = append(currnode.handlers, handlers...)
	return nil
}

func (n *node) solve(relativePath string) ([]Handler, map[string]string, error) {
	currnode := n
	segs := strings.Split(relativePath, "/")
	params := make(map[string]string)

	// If the size of segs is less than 2, the given relative path does not start
	// "/", so the path is invalid.
	if len(segs) <= 1 {
		return nil, nil, ErrInvalidPath
	}

	// If the size of segs is 2 and segs[1] is empty string, the given relative
	// path is "/", so return handlers of n.
	if len(segs) == 2 && segs[1] == "" {
		if len(currnode.handlers) == 0 {
			return nil, nil, ErrNotExistPath
		}
		return currnode.handlers, nil, nil
	}

	for i := 1; i < len(segs); i++ {
		seg := segs[i]
		for _, c := range currnode.children {
			// /user/:name will match /user/john, but not match /user/, /user.
			// /user/:name/*action will match /user/john/send and /user/john/, not /user/john.
			switch c.name[0] {
			case ':':
				ss := c.name[1:len(c.name)]
				// If seg is an empty string, the end of given relative path is "/",
				// so wildcard ':' doesn't match it.
				if seg == "" {
					return nil, nil, ErrInvalidParameterPath
				}
				params[ss] = seg
				currnode = c
				goto exist_segment
			case '*':
				ss := c.name[1:len(c.name)]
				params[ss] = seg
				currnode = c
				goto exist_segment
			default:
				if c.name == seg {
					currnode = c
					goto exist_segment
				}
			}
		}
		return nil, nil, ErrNotExistPath
	exist_segment:
	}
	if len(currnode.handlers) == 0 {
		return nil, nil, ErrNotExistPath
	}
	return currnode.handlers, params, nil
}
