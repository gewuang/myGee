package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	w      http.ResponseWriter
	req    *http.Request
	method string
	path   string
	code   int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		w:      w,
		req:    req,
		method: req.Method,
		path:   req.URL.Path,
	}
}

// 从url中取值
func (c *Context) PostFrom(key string) string {
	return c.req.PostFormValue(key)
}

// 从url中取值
func (c *Context) Query(key string) string {
	return c.req.URL.Query().Get(key)
}

// 设置头部
func (c *Context) SetHeader(key, value string) {
	c.w.Header().Set(key, value)
}

// 设置code
func (c *Context) SetCode(code int) {
	c.code = code
	c.w.WriteHeader(code)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetCode(code)
	c.w.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Data(code int, data []byte) {
	c.SetCode(code)
	c.w.Write([]byte(data))
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetCode(code)
	c.w.Write([]byte(html))
}
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetCode(code)
	encoder := json.NewEncoder(c.w)

	if err := encoder.Encode(obj); err != nil {
		http.Error(c.w, err.Error(), 500)
	}
}
