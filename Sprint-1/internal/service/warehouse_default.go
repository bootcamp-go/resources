package service

import "github.com/usuario/repositorio/internal"

// NewWarehouseDefault creates a new instance of the warehouse service
func NewWarehouseDefault(rp internal.WarehouseRepository) *WarehouseDefault {
	return &WarehouseDefault{
		rp: rp,
	}
}

// WarehouseDefault is the default implementation of the warehouse service
type WarehouseDefault struct {
	// rp is the repository used by the service
	rp internal.WarehouseRepository
}

// FindAll returns all warehouses
func (s *WarehouseDefault) FindAll() (warehouses []internal.Warehouse, err error) {
	return
}

// FindByID returns a warehouse
func (s *WarehouseDefault) FindByID(id int) (warehouse internal.Warehouse, err error) {
	return
}

// Save creates a new warehouse
func (s *WarehouseDefault) Save(warehouse *internal.Warehouse) (err error) {
	return
}

// Update updates a warehouse
func (s *WarehouseDefault) Update(warehouse *internal.Warehouse) (err error) {
	return
}

// Delete deletes a warehouse
func (s *WarehouseDefault) Delete(id int) (err error) {
	return
}