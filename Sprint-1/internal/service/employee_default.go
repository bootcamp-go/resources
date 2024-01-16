package service

import "github.com/usuario/repositorio/internal"

// NewEmployeeDefault creates a new instance of the employee service
func NewEmployeeDefault(rp internal.EmployeeRepository) *EmployeeDefault {
	return &EmployeeDefault{
		rp: rp,
	}
}

// EmployeeDefault is the default implementation of the employee service
type EmployeeDefault struct {
	// rp is the repository used by the service
	rp internal.EmployeeRepository
}

// FindAll returns all employees
func (s *EmployeeDefault) FindAll() (employees []internal.Employee, err error) {
	return
}

// FindByID returns a employee
func (s *EmployeeDefault) FindByID(id int) (employee internal.Employee, err error) {
	return
}

// Save creates a new employee
func (s *EmployeeDefault) Save(employee *internal.Employee) (err error) {
	return
}

// Update updates a employee
func (s *EmployeeDefault) Update(employee *internal.Employee) (err error) {
	return
}

// Delete deletes a employee
func (s *EmployeeDefault) Delete(id int) (err error) {
	return
}