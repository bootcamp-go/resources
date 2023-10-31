package repository

import "app/internal"

// NewRepositoryReadVehicleMap is a function that returns a new instance of RepositoryReadVehicleMap
func NewRepositoryReadVehicleMap(db map[int]internal.Vehicle) *RepositoryReadVehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &RepositoryReadVehicleMap{db: defaultDb}
}

// RepositoryReadVehicleMap is a struct that represents a vehicle repository
type RepositoryReadVehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *RepositoryReadVehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// FindByColorAndYear is a method that returns a map of vehicles that match the color and fabrication year
func (r *RepositoryReadVehicleMap) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// filter db
	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == fabricationYear {
			v[key] = value
		}
	}

	return
}

// FindByBrandAndYearRange is a method that returns a map of vehicles that match the brand and a range of fabrication years
func (r *RepositoryReadVehicleMap) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// filter db
	for key, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= startYear && value.FabricationYear <= endYear {
			v[key] = value
		}
	}

	return
}

// FindByBrand is a method that returns a map of vehicles that match the brand
func (r *RepositoryReadVehicleMap) FindByBrand(brand string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// filter db
	for key, value := range r.db {
		if value.Brand == brand {
			v[key] = value
		}
	}

	return
}

// FindByWeightRange is a method that returns a map of vehicles that match the weight range
func (r *RepositoryReadVehicleMap) FindByWeightRange(fromWeight float64, toWeight float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// filter db
	for key, value := range r.db {
		if value.Weight >= fromWeight && value.Weight <= toWeight {
			v[key] = value
		}
	}

	return
}