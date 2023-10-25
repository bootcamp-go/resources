package application

import (
	"app/internal/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigApplicationDefault is a struct that represents the configuration for ApplicationDefault
type ConfigApplicationDefault struct {
	// Router is the router / multiplexer that will be used by the application
	Router *chi.Mux
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
	LoaderFilePath string
}

// NewApplicationDefault is a function that returns a new instance of ApplicationDefault
func NewApplicationDefault(cfg *ConfigApplicationDefault) *ApplicationDefault {
	// default values
	defaultConfig := &ConfigApplicationDefault{
		Router: chi.NewRouter(),
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.Router != nil {
			defaultConfig.Router = cfg.Router
		}
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ApplicationDefault{
		router: defaultConfig.Router,
		serverAddress: defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ApplicationDefault is a struct that implements the Application interface
type ApplicationDefault struct {
	// router is the router / multiplexer that will be used by the application
	router *chi.Mux
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// SetUp is a method that sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - loader: loader for vehicles
	ld := loader.NewLoaderVehicleJSON(a.loaderFilePath)
	// - db: map of vehicles
	db, err := ld.Load()
	if err != nil {
		return
	}
	// - repository: repository for vehicles
	rp := repository.NewRepositoryReadVehicleMap(db)
	// - service: service for vehicles
	sv := service.NewServiceVehicleDefault(rp)
	// - handler: handler for vehicles
	hd := handler.NewHandlerVehicle(sv)

	// routes
	// - middlewares
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)
	// - endpoints
	a.router.Route("/vehicles", func(r chi.Router) {
		// Get vehicles by color and year
		r.Get("/color/{color}/year/{year}", hd.FindByColorAndYear())
		// Get vehicles by brand between years
		r.Get("/brand/{brand}/between/{start_year}/{end_year}", hd.FindByBrandAndYearRange())
		// Get average max speed by brand
		r.Get("/average_speed/brand/{brand}", hd.AverageMaxSpeedByBrand())
		// Get average capacity by brand
		r.Get("/average_capacity/brand/{brand}", hd.AverageCapacityByBrand())
		// Get vehicles by weight range (query)
		r.Get("/weight", hd.SearchByWeightRange())
	})

	return
}

// Run is a method that runs the application
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddress, a.router)
	return
}