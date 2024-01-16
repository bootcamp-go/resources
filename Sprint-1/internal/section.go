package internal

import "errors"

// Section is a struct that contains the section's information
type Section struct {
	// ID is the unique identifier of the section
	ID int
	// SectionNumber is the number of the section
	SectionNumber int
	// CurrentTemperature is the current temperature of the section
	CurrentTemperature float64
	// MinimumTemperature is the minimum temperature that can be maintained in the section
	MinimumTemperature float64
	// CurrentCapacity is the current capacity of the section
	CurrentCapacity int
	// MinimumCapacity is the minimum capacity of the section
	MinimumCapacity int
	// MaximumCapacity is the maximum capacity of the section
	MaximumCapacity int
	// WarehouseID is the unique identifier of the warehouse to which the section belongs
	WarehouseID int
	// ProductTypeID is the unique identifier of the type of product stored in the section
	ProductTypeID int
}

var (
	// ErrSectionRepositoryNotFound is returned when the section is not found
	ErrSectionRepositoryNotFound = errors.New("repository: section not found")
	// ErrSectionRepositoryDuplicated is returned when the section already exists
	ErrSectionRepositoryDuplicated = errors.New("repository: section already exists")
)

// SectionRepository is an interface that contains the methods that the section repository should support
type SectionRepository interface {
	// FindAll returns all the sections
	FindAll() ([]Section, error)
	// FindByID returns the section with the given ID
	FindByID(id int) (Section, error)
	// Save saves the given section
	Save(section *Section) error
	// Update updates the given section
	Update(section *Section) error
	// Delete deletes the section with the given ID
	Delete(id int) error
}

// SectionService is an interface that contains the methods that the section service should support
type SectionService interface {
	// FindAll returns all the sections
	FindAll() ([]Section, error)
	// FindByID returns the section with the given ID
	FindByID(id int) (Section, error)
	// Save saves the given section
	Save(section *Section) error
	// Update updates the given section
	Update(section *Section) error
	// Delete deletes the section with the given ID
	Delete(id int) error
}