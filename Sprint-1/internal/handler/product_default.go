package handler

import (
	"net/http"

	"github.com/usuario/repositorio/internal"
)

// NewProductDefault creates a new instance of the product handler
func NewProductDefault(sv internal.ProductService) *ProductDefault {
	return &ProductDefault{
		sv: sv,
	}
}

// ProductDefault is the default implementation of the product handler
type ProductDefault struct {
	// sv is the service used by the handler
	sv internal.ProductService
}

// GetAll returns all products
func (h *ProductDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a product
func (h *ProductDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new product
func (h *ProductDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a product
func (h *ProductDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a product
func (h *ProductDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}