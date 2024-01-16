package handler

import (
	"net/http"

	"github.com/usuario/repositorio/internal"
)

// NewSectionDefault creates a new instance of the section handler
func NewSectionDefault(sv internal.SectionService) *SectionDefault {
	return &SectionDefault{
		sv: sv,
	}
}

// SectionDefault is the default implementation of the section handler
type SectionDefault struct {
	// sv is the service used by the handler
	sv internal.SectionService
}

// GetAll returns all sections
func (h *SectionDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a section
func (h *SectionDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new section
func (h *SectionDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a section
func (h *SectionDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a section
func (h *SectionDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}