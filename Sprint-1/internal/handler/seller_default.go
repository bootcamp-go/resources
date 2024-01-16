package handler

import (
	"net/http"

	"github.com/usuario/repositorio/internal"
)

// NewSellerDefault creates a new instance of the seller handler
func NewSellerDefault(sv internal.SellerService) *SellerDefault {
	return &SellerDefault{
		sv: sv,
	}
}

// SellerDefault is the default implementation of the seller handler
type SellerDefault struct {
	// sv is the service used by the handler
	sv internal.SellerService
}

// GetAll returns all sellers
func (h *SellerDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a seller
func (h *SellerDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new seller
func (h *SellerDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a seller
func (h *SellerDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a seller
func (h *SellerDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}