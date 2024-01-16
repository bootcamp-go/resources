package service

import "github.com/usuario/repositorio/internal"

// NewProductDefault creates a new instance of the product service
func NewProductDefault(rp internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}
}

// ProductDefault is the default implementation of the product service
type ProductDefault struct {
	// rp is the repository used by the service
	rp internal.ProductRepository
}

// FindAll returns all products
func (s *ProductDefault) FindAll() (products []internal.Product, err error) {
	return
}

// FindByID returns a product
func (s *ProductDefault) FindByID(id int) (product internal.Product, err error) {
	return
}

// Save creates a new product
func (s *ProductDefault) Save(product *internal.Product) (err error) {
	return
}

// Update updates a product
func (s *ProductDefault) Update(product *internal.Product) (err error) {
	return
}

// Delete deletes a product
func (s *ProductDefault) Delete(id int) (err error) {
	return
}