package internal

import "errors"

// Buyer is a struct that contains the buyer's information
type Buyer struct {
	// ID is the unique identifier of the buyer
	ID int
	// CardNumberID is the unique identifier of the card number
	CardNumberID int
	// FirstName is the first name of the buyer
	FirstName string
	// LastName is the last name of the buyer
	LastName string
}

var (
	// ErrBuyerRepositoryNotFound is returned when the buyer is not found
	ErrBuyerRepositoryNotFound = errors.New("repository: buyer not found")
	// ErrBuyerRepositoryDuplicated is returned when the buyer already exists
	ErrBuyerRepositoryDuplicated = errors.New("repository: buyer already exists")
)

// BuyerRepository is an interface that contains the methods that the buyer repository should support
type BuyerRepository interface {
	// FindAll returns all the buyers
	FindAll() ([]Buyer, error)
	// FindByID returns the buyer with the given ID
	FindByID(id int) (Buyer, error)
	// Save saves the given buyer
	Save(buyer *Buyer) error
	// Update updates the given buyer
	Update(buyer *Buyer) error
	// Delete deletes the buyer with the given ID
	Delete(id int) error
}

// BuyerService is an interface that contains the methods that the buyer service should support
type BuyerService interface {
	// FindAll returns all the buyers
	FindAll() ([]Buyer, error)
	// FindByID returns the buyer with the given ID
	FindByID(id int) (Buyer, error)
	// Save saves the given buyer
	Save(buyer *Buyer) error
	// Update updates the given buyer
	Update(buyer *Buyer) error
	// Delete deletes the buyer with the given ID
	Delete(id int) error
}