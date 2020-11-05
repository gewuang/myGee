package gee

import (
	"log"
	"net/http"
)

type HandleFunc func(c *Context)

type engine struct {
	r *router
}

func (e *engine) addRouter(method, path string, handle HandleFunc) {
	e.r.addRouter(method, path, handle)
}

func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.r.handler(c)
}

func NEW() *engine {
	return &engine{r: newRouter()}
}

func (e *engine) GET(path string, handle HandleFunc) {
	e.addRouter("GET", path, handle)
}
func (e *engine) POST(path string, handle HandleFunc) {
	e.addRouter("POST", path, handle)
}
func (e *engine) DELETE(path string, handle HandleFunc) {
	e.addRouter("DELETE", path, handle)
}
func (e *engine) RUN() {
	log.Fatal(http.ListenAndServe("localhost:8989", e))
}
