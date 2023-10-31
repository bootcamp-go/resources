package internal

import "errors"

var (
	// ErrServiceInvalidFind is an error that represents an invalid find
	ErrServiceInvalidFind = errors.New("service: invalid find")
	// ErrServiceInvalidSearch is an error that represents an invalid search
	ErrServiceInvalidSearch = errors.New("service: invalid search")
	// ErrServiceNoVehicles is an error that represents no vehicles
	ErrServiceNoVehicles = errors.New("service: no vehicles")
)

// SearchQuery is a struct that represents a search query
type SearchQuery struct {
	// FromWeight is the minimum weight
	FromWeight float64
	// ToWeight is the maximum weight
	ToWeight float64
}

// ServiceVehicle is an interface that represents a vehicle service
type ServiceVehicle interface {
	// FindByColorAndYear is a method that returns a map of vehicles that match the color and fabrication year
	FindByColorAndYear(color string, fabricationYear int) (v map[int]Vehicle, err error)

	// FindByBrandAndYearRange is a method that returns a map of vehicles that match the brand and a range of fabrication years
	FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]Vehicle, err error)

	// AverageMaxSpeedByBrand is a method that returns the average speed of the vehicles by brand
	AverageMaxSpeedByBrand(brand string) (a float64, err error)

	// AverageCapacityByBrand is a method that returns the average capacity of the vehicles by brand
	AverageCapacityByBrand(brand string) (a int, err error)

	// SearchByWeightRange
	// - method: hybrid. usage of static procedure and static optional (not dynamic types such as maps or slices)
	// - query:
	// 	 !ok -> will return all vehicles
	// 	 ok  -> will return filtered vehicles
	SearchByWeightRange(query SearchQuery, ok bool) (v map[int]Vehicle, err error)
}