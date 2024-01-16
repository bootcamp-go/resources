package internal

import "errors"

// Employee is a struct that contains the employee's information
type Employee struct {
	// ID is the unique identifier of the employee
	ID int
	// CardNumberID is the unique identifier of the card number
	CardNumberID int
	// FirstName is the first name of the employee
	FirstName string
	// LastName is the last name of the employee
	LastName string
	// WarehouseID is the unique identifier of the warehouse to which the employee belongs
	WarehouseID int
}

var (
	// ErrEmployeeRepositoryNotFound is returned when the employee is not found
	ErrEmployeeRepositoryNotFound = errors.New("repository: employee not found")
	// ErrEmployeeRepositoryDuplicated is returned when the employee already exists
	ErrEmployeeRepositoryDuplicated = errors.New("repository: employee already exists")
)

// EmployeeRepository is an interface that contains the methods that the employee repository should support
type EmployeeRepository interface {
	// FindAll returns all the employees
	FindAll() ([]Employee, error)
	// FindByID returns the employee with the given ID
	FindByID(id int) (Employee, error)
	// Save saves the given employee
	Save(employee *Employee) error
	// Update updates the given employee
	Update(employee *Employee) error
	// Delete deletes the employee with the given ID
	Delete(id int) error
}

// EmployeeService is an interface that contains the methods that the employee service should support
type EmployeeService interface {
	// FindAll returns all the employees
	FindAll() ([]Employee, error)
	// FindByID returns the employee with the given ID
	FindByID(id int) (Employee, error)
	// Save saves the given employee
	Save(employee *Employee) error
	// Update updates the given employee
	Update(employee *Employee) error
	// Delete deletes the employee with the given ID
	Delete(id int) error
}