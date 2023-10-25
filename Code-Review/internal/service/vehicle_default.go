package service

import "app/internal"

// ServiceVehicleDefault is a struct that represents the default service for vehicles
type ServiceVehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.RepositoryReadVehicle
}

// NewServiceVehicleDefault is a function that returns a new instance of ServiceVehicleDefault
func NewServiceVehicleDefault(rp internal.RepositoryReadVehicle) *ServiceVehicleDefault {
	return &ServiceVehicleDefault{rp: rp}
}

// FindByColorAndYear is a method that returns a map of vehicles that match the color and fabrication year
func (s *ServiceVehicleDefault) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByColorAndYear(color, fabricationYear)
	return
}

// FindByBrandAndYearRange is a method that returns a map of vehicles that match the brand and a range of fabrication years
func (s *ServiceVehicleDefault) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByBrandAndYearRange(brand, startYear, endYear)
	return
}

// AverageMaxSpeedByBrand is a method that returns the average speed of the vehicles by brand
func (s *ServiceVehicleDefault) AverageMaxSpeedByBrand(brand string) (a float64, err error) {
	// get vehicles by brand
	v, err := s.rp.FindByBrand(brand)
	if err != nil {
		return
	}

	// check if there are vehicles
	if len(v) == 0 {
		err = internal.ErrServiceNoVehicles
		return
	}

	var totalSpeed float64
	for _, vehicle := range v {
		totalSpeed += vehicle.MaxSpeed
	}

	a = totalSpeed / float64(len(v))
	return
}
		
// AverageCapacityByBrand is a method that returns the average capacity of the vehicles by brand
func (s *ServiceVehicleDefault) AverageCapacityByBrand(brand string) (a int, err error) {
	// get vehicles by brand
	v, err := s.rp.FindByBrand(brand)
	if err != nil {
		return
	}
	
	// check if there are vehicles
	if len(v) == 0 {
		err = internal.ErrServiceNoVehicles
		return
	}

	var totalCapacity int
	for _, vehicle := range v {
		totalCapacity += vehicle.Capacity
	}

	a = totalCapacity / len(v)
	return
}

// SearchByWeightRange
func (s *ServiceVehicleDefault) SearchByWeightRange(query internal.SearchQuery, ok bool) (v map[int]internal.Vehicle, err error) {
	// check if query is set
	if !ok {
		v, err = s.rp.FindAll()
		return
	}

	v, err = s.rp.FindByWeightRange(query.FromWeight, query.ToWeight)
	return
}
	