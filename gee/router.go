package gee

import (
	"fmt"
)

type router struct {
	handle map[string]HandleFunc
}

func newRouter() *router {
	return &router{handle: make(map[string]HandleFunc)}
}

func (r *router) addRouter(method, path string, handle HandleFunc) {
	str := method + "-" + path
	r.handle[str] = handle
}

func (r *router) handler(c *Context) {
	key := c.req.Method + "-" + c.req.URL.Path

	if handle, ok := r.handle[key]; ok {
		handle(c)
	} else {
		fmt.Fprintf(c.w, "404 not found\n")
	}
}
