package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// HandlerFunc is the type of the handler functions
type HandlerFunc func(w http.ResponseWriter, r *http.Request) (err error)

// RouterGroup is an struct that abstracts the router of chi
type RouterGroup struct {
	// rt is the chi router
	rt *chi.Mux
	// md is the middleware stack
	md []func(HandlerFunc) HandlerFunc
	// basePath is the basePath of the router
	basePath string
}

// Use adds a middleware to the middleware stack
func (rg *RouterGroup) Use(md ...func(HandlerFunc) HandlerFunc) {
	(*rg).md = append((*rg).md, md...)
}

// Handle adds a route with the specified methods
func (rg *RouterGroup) Handle(method string, path string, hd HandlerFunc, md ...func(HandlerFunc) HandlerFunc) {
	// chain
	// - handler with post middlewares
	hd = handlerChain(hd, md...)
	// - handler with group middlewares
	hd = handlerChain(hd, (*rg).md...)

	// handler and path
	handler := handlerAdapter(hd)
	path = (*rg).basePath + path

	// register the route
	(*rg).rt.Method(method, path, handler)
}

// Route allows to work with nested routers
// - dynamic: fn allows to work with a group (it contains previous path and middlewares)
// notes: same idea can be done with a func SubRoute returning previous path and middlewares and working directly from the instance
func (rg *RouterGroup) Route(path string, fn func (rg *RouterGroup)) {
	// new sub router
	subRouter := &RouterGroup{
		rt: (*rg).rt,
		md: (*rg).md,
		basePath: (*rg).basePath + path,
	}

	// call the function
	fn(subRouter)
}


/*
	tool functions
	- chain: builds the decorated handler
	- handlerAdapter: adapts the handler to the native http handler
*/
// handlerChain builds the decorated handler
func handlerChain(hd HandlerFunc, md ...func(HandlerFunc) HandlerFunc) HandlerFunc {
	// chain handler and middlewares
	for _, m := range md {
		hd = m(hd)
	}
	return hd
}
// handlerAdapter adapts the handler to the native http handler
func handlerAdapter(hd HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		err := hd(w, req)
		if err != nil {
			panic(err)
		}
	}
}