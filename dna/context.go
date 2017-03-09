package dna

import (
	"net/http"
)

var (
	// Gin のそれに倣い, ハンドラチェーンを途中で切る場合は,
	// ハンドラ内で Abort し, return するという方針を取る.
	// ServeHTTP 側で Context の index が abortIndex を超えるか
	// をチェックし, 超える場合は処理を中断する方針を取る.
	abortIndex = 10
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request

	// URLParams represents parameters of URL.
	URLParams map[string]string

	// index represents an index of the handler chain.
	index int

	Any map[string]interface{}
}

func NewContext(w http.ResponseWriter, r *http.Request, params map[string]string) *Context {
	return &Context{
		Writer:    w,
		Request:   r,
		URLParams: params,
		index:     0,
	}
}

func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) AbortWithStatus(code int) {
	c.Status(code)
	c.Abort()
}

func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)
}

// JSON renders JSON object of v to c.Writer.
func (c *Context) JSON(v interface{}) {
	if err := EncodeJson(v, c.Writer); err != nil {
		c.Status(500)
	}
}
