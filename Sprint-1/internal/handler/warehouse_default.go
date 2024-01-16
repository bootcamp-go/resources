package handler

import (
	"net/http"

	"github.com/usuario/repositorio/internal"
)

// NewWarehouseDefault creates a new instance of the warehouse handler
func NewWarehouseDefault(sv internal.WarehouseService) *WarehouseDefault {
	return &WarehouseDefault{
		sv: sv,
	}
}

// WarehouseDefault is the default implementation of the warehouse handler
type WarehouseDefault struct {
	// sv is the service used by the handler
	sv internal.WarehouseService
}

// GetAll returns all warehouses
func (h *WarehouseDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a warehouse
func (h *WarehouseDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new warehouse
func (h *WarehouseDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a warehouse
func (h *WarehouseDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a warehouse
func (h *WarehouseDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}