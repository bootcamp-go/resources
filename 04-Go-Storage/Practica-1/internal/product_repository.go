package internal

import "errors"

var (
	// ErrRepositoryProductNotFound is returned when a product is not found.
	ErrRepositoryProductNotFound = errors.New("repository: product not found")
)

// RepositoryProduct is an interface that contains the methods for a product repository
type RepositoryProduct interface {
	// FindById returns a product by its id
	FindById(id int) (p Product, err error)
	// Save saves a product
	Save(p *Product) (err error)
	// UpdateOrSave updates or saves a product
	UpdateOrSave(p *Product) (err error)
	// Update updates a product
	Update(p *Product) (err error)
	// Delete deletes a product
	Delete(id int) (err error)
}