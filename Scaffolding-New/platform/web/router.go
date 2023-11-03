package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewRouter creates a new router
func NewRouter() *Router {
	return &Router{
		routerGroup: RouterGroup{
			rt: chi.NewRouter(),
		},
	}
}

// Router is an struct that abstracts the router of chi
type Router struct {
	// routerGroup is the router group
	routerGroup RouterGroup
}

// Use adds a middleware to the middleware stack
func (r *Router) Use(md ...func(HandlerFunc) HandlerFunc) {
	r.routerGroup.Use(md...)
}

// Handle adds a route with the specified methods
func (r *Router) Handle(method string, path string, hd HandlerFunc, md ...func(HandlerFunc) HandlerFunc) {
	r.routerGroup.Handle(method, path, hd, md...)
}

// Route allows to work with nested routers
// - dynamic: fn allows to work with a group (it contains previous path and middlewares)
// notes: same idea can be done with a func SubRoute returning previous path and middlewares and working directly from the instance
func (r *Router) Route(path string, fn func (rg *RouterGroup)) {
	r.routerGroup.Route(path, fn)
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.routerGroup.rt.ServeHTTP(w, req)
}

// Run starts the server
func (r *Router) Run(addr string) (err error) {
	err = http.ListenAndServe(addr, r)
	return
}