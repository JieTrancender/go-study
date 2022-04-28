package kbm

import (
	"net/http"
)

// HandlerFunc defines the request handler used by kbm
type HandlerFunc func(*Context)

// Engine implements the interface of ServeHttp
type Engine struct {
	router *router
}

// New is the constructor of kbm.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRouter(method, pattern string, handler HandlerFunc) {
	engine.router.addRouter(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRouter("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRouter("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
