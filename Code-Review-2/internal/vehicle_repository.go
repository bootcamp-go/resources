package internal

import "errors"

var (
	// ErrRepositoryInvalidFind is an error that represents an invalid find
	ErrRepositoryInvalidFind = errors.New("repository: invalid find")
)

// RepositoryReadVehicle is an interface that represents a vehicle repository
// - method: static. All searchs are strong typed, not hybrid or dynamic
type RepositoryReadVehicle interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)

	// FindByColorAndYear is a method that returns a map of vehicles that match the color and fabrication year
	FindByColorAndYear(color string, fabricationYear int) (v map[int]Vehicle, err error)

	// FindByBrandAndYearRange is a method that returns a map of vehicles that match the brand and a range of fabrication years
	FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]Vehicle, err error)

	// FindByBrand is a method that returns a map of vehicles that match the brand
	FindByBrand(brand string) (v map[int]Vehicle, err error)

	// FindByWeightRange is a method that returns a map of vehicles that match the weight range
	FindByWeightRange(fromWeight float64, toWeight float64) (v map[int]Vehicle, err error)
}