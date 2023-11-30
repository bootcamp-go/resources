package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigApplicationDefault is the configuration of the application.
type ConfigApplicationDefault struct {
	// Addr is the address of the application.
	Addr string
}

// NewApplicationDefault returns a new ApplicationDefault.
func NewApplicationDefault(cfg *ConfigApplicationDefault) *ApplicationDefault {
	// default values
	defaultRt  := chi.NewRouter()
	defaultCfg := &ConfigApplicationDefault{
		Addr: ":8080",
	}
	if cfg != nil {
		if cfg.Addr != "" {
			defaultCfg.Addr = cfg.Addr
		}
	}

	return &ApplicationDefault{
		rt:   defaultRt,
		addr: defaultCfg.Addr,
	}
}

// ApplicationDefault is an struct that implements the Application interface.
type ApplicationDefault struct {
	// rt is the router of the application.
	rt *chi.Mux
	// addr is the address of the application.
	addr string
}

// TearDown tears down the application.
// - should be used as a defer function
func (a *ApplicationDefault) TearDown() (err error) {
	return
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - repository
	rpProduct := repository.NewProductsMap(nil)
	// - handler
	hdProduct := handler.NewProductsDefault(rpProduct)

	// router
	// - middleware
	a.rt.Use(middleware.Logger)
	a.rt.Use(middleware.Recoverer)
	// - endpoints
	a.rt.Route("/product", func(r chi.Router) {
		// - GET /product
		r.Get("/", hdProduct.Get())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.addr, a.rt)
	return
}