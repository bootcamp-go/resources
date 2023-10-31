package internal

// LoaderVehicle is an interface that represents the loader for vehicles
type LoaderVehicle interface {
	// Load is a method that loads the vehicles
	Load() (v map[int]Vehicle, err error)
}