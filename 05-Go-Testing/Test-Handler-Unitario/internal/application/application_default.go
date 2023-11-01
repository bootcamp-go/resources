package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewDefault returns a new Default
func NewDefault(addr string) (d *Default) {
	// default config
	defaultRouter := chi.NewRouter()
	defaultAddr := ":8080"
	if addr != "" {
		defaultAddr = addr
	}

	d = &Default{
		rt:   defaultRouter,
		addr: defaultAddr,
	}
	return
}

// Default is a struct that contains the default application
type Default struct {
	// rt is the router of the server
	rt *chi.Mux
	// addr is the address of the server
	addr string
}

// TearDown tears down the default application
func (d *Default) TearDown() (err error) {
	return
}

// SetUp sets up the default application
func (d *Default) SetUp() (err error) {
	// dependencies
	// - repository: map
	rp := repository.NewProductMap(nil, 0)
	// - handler
	hd := handler.NewProduct(rp)

	// router
	// - middlewares
	d.rt.Use(middleware.Logger)
	d.rt.Use(middleware.Recoverer)
	// - routes / endpoints
	d.rt.Post("/product", hd.Create())

	return
}

// Run runs the default application
func (d *Default) Run() (err error) {
	err = http.ListenAndServe(d.addr, d.rt)
	return
}