package internal

import "errors"

// Warehouse is a struct that contains the warehouse's information
type Warehouse struct {
	// ID is the unique identifier of the warehouse
	ID int
	// WarehouseCode is the unique code of the warehouse
	WarehouseCode string
	// Address is the address of the warehouse
	Address string
	// Telephone is the telephone number of the warehouse
	Telephone string
	// MinimumCapacity is the minimum capacity of the warehouse
	MinimumCapacity int
	// MinimumTemperature is the minimum temperature that can be maintained in the warehouse
	MinimumTemperature float64
}

var (
	// ErrWarehouseRepositoryNotFound is returned when the warehouse is not found
	ErrWarehouseRepositoryNotFound = errors.New("repository: warehouse not found")
	// ErrWarehouseRepositoryDuplicated is returned when the warehouse already exists
	ErrWarehouseRepositoryDuplicated = errors.New("repository: warehouse already exists")
)

// WarehouseRepository is an interface that contains the methods that the warehouse repository should support
type WarehouseRepository interface {
	// FindAll returns all the warehouses
	FindAll() ([]Warehouse, error)
	// FindByID returns the warehouse with the given ID
	FindByID(id int) (Warehouse, error)
	// Save saves the given warehouse
	Save(warehouse *Warehouse) error
	// Update updates the given warehouse
	Update(warehouse *Warehouse) error
	// Delete deletes the warehouse with the given ID
	Delete(id int) error
}

// WarehouseService is an interface that contains the methods that the warehouse service should support
type WarehouseService interface {
	// FindAll returns all the warehouses
	FindAll() ([]Warehouse, error)
	// FindByID returns the warehouse with the given ID
	FindByID(id int) (Warehouse, error)
	// Save saves the given warehouse
	Save(warehouse *Warehouse) error
	// Update updates the given warehouse
	Update(warehouse *Warehouse) error
	// Delete deletes the warehouse with the given ID
	Delete(id int) error
}
