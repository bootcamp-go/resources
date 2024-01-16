package application

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigServerChi is the configuration for the server
type ConfigServerChi struct {
	// Addr is the address to listen on
	Addr string
	// MySQLDSN is the DSN for the MySQL database
	MySQLDSN string
}

// NewServerChi creates a new instance of the server
func NewServerChi(cfg ConfigServerChi) *ServerChi {
	// default config
	defaultCfg := ConfigServerChi{
		Addr:     ":8080",
		MySQLDSN: "",
	}
	if cfg.Addr != "" {
		defaultCfg.Addr = cfg.Addr
	}
	if cfg.MySQLDSN != "" {
		defaultCfg.MySQLDSN = cfg.MySQLDSN
	}

	return &ServerChi{
		addr:     defaultCfg.Addr,
		mysqlDSN: defaultCfg.MySQLDSN,
	}
}

// ServerChi is the default implementation of the server
type ServerChi struct {
	// addr is the address to listen on
	addr string
	// mysqlDSN is the DSN for the MySQL database
	mysqlDSN string
}

// Run runs the server
func (s *ServerChi) Run() (err error) {
	// dependencies
	// - database: connection
	db, err := sql.Open("mysql", s.mysqlDSN)
	if err != nil {
		return
	}
	defer db.Close()
	// - database: ping
	err = db.Ping()
	if err != nil {
		return
	}

	// - router
	router := chi.NewRouter()
	//   middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	//   endpoints
	//     sellers
	buildSellersRouter(router)
	//     warehouses
	buildWarehousesRouter(router)
	//     sections
	buildSectionsRouter(router)
	//     products
	buildProductsRouter(router)
	//     employees
	buildEmployeesRouter(router)
	//     buyers
	buildBuyersRouter(router)

	// run
	err = http.ListenAndServe(s.addr, router)
	return
}

// buildSellersRouter builds the router for the sellers endpoints
func buildSellersRouter(router *chi.Mux) {
	// ...
}

// buildWarehousesRouter builds the router for the warehouses endpoints
func buildWarehousesRouter(router *chi.Mux) {
	// ...
}

// buildSectionsRouter builds the router for the sections endpoints
func buildSectionsRouter(router *chi.Mux) {
	// ...
}

// buildProductsRouter builds the router for the products endpoints
func buildProductsRouter(router *chi.Mux) {
	// ...
}

// buildEmployeesRouter builds the router for the employees endpoints
func buildEmployeesRouter(router *chi.Mux) {
	// ...
}

// buildBuyersRouter builds the router for the buyers endpoints
func buildBuyersRouter(router *chi.Mux) {
	// ...
}