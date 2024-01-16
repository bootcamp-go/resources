package handler

import (
	"net/http"

	"github.com/usuario/repositorio/internal"
)

// NewEmployeeDefault creates a new instance of the employee handler
func NewEmployeeDefault(sv internal.EmployeeService) *EmployeeDefault {
	return &EmployeeDefault{
		sv: sv,
	}
}

// EmployeeDefault is the default implementation of the employee handler
type EmployeeDefault struct {
	// sv is the service used by the handler
	sv internal.EmployeeService
}

// GetAll returns all employees
func (h *EmployeeDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a employee
func (h *EmployeeDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new employee
func (h *EmployeeDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a employee
func (h *EmployeeDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a employee
func (h *EmployeeDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}