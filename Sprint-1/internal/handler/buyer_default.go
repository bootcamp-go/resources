package handler

import (
	"net/http"

	"github.com/usuario/repositorio/internal"
)

// NewBuyerDefault creates a new instance of the buyer handler
func NewBuyerDefault(sv internal.BuyerService) *BuyerDefault {
	return &BuyerDefault{
		sv: sv,
	}
}

// BuyerDefault is the default implementation of the buyer handler
type BuyerDefault struct {
	// sv is the service used by the handler
	sv internal.BuyerService
}

// GetAll returns all buyers
func (h *BuyerDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a buyer
func (h *BuyerDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new buyer
func (h *BuyerDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a buyer
func (h *BuyerDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a buyer
func (h *BuyerDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}