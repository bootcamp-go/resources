package application

import (
	"app/internal/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"

	"github.com/gin-gonic/gin"
)

// ConfigApplicationDefault is a struct that represents the configuration for ApplicationDefault
type ConfigApplicationDefault struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
	LoaderFilePath string
}

// NewApplicationDefault is a function that returns a new instance of ApplicationDefault
func NewApplicationDefault(cfg *ConfigApplicationDefault) *ApplicationDefault {
	// default values
	defaultRouter := gin.New()
	defaultConfig := &ConfigApplicationDefault{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ApplicationDefault{
		router: defaultRouter,
		serverAddress: defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ApplicationDefault is a struct that implements the Application interface
type ApplicationDefault struct {
	// router is the router / multiplexer that will be used by the application
	router *gin.Engine
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
	a.router.Use(gin.Logger())
	a.router.Use(gin.Recovery())
	// - endpoints
	grVehicles := a.router.Group("/vehicles")
	// Get vehicles by color and year
	grVehicles.GET("/color/:color/year/:year", hd.FindByColorAndYear())
	// Get vehicles by brand between years
	grVehicles.GET("/brand/:brand/between/:start_year/:end_year", hd.FindByBrandAndYearRange())
	// Get average max speed by brand
	grVehicles.GET("/average_speed/brand/:brand", hd.AverageMaxSpeedByBrand())
	// Get average capacity by brand
	grVehicles.GET("/average_capacity/brand/:brand", hd.AverageCapacityByBrand())
	// Get vehicles by weight range (query)
	grVehicles.GET("/weight", hd.SearchByWeightRange())

	return
}

// Run is a method that runs the application
func (a *ApplicationDefault) Run() (err error) {
	err = a.router.Run(a.serverAddress)
	return
}