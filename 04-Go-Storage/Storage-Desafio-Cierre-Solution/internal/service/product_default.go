package service

import "app/internal"

// NewProductsDefault creates new default service for product entity.
func NewProductsDefault(rp internal.RepositoryProduct) *ProducstDefault {
	return &ProducstDefault{rp}
}

// ProducstDefault is the default service implementation for product entity.
type ProducstDefault struct {
	// rp is the repository for product entity.
	rp internal.RepositoryProduct
}

// FindAll returns all products.
func (s *ProducstDefault) FindAll() (p []internal.Product, err error) {
	p, err = s.rp.FindAll()
	return
}

// FindTopProductsByAmountSold returns the top products by amount sold.
func (s *ProducstDefault) FindTopProductsByAmountSold(limit int) (p []internal.ProductAmountSold, err error) {
	p, err = s.rp.FindTopProductsByAmountSold(limit)
	return
}

// Save saves the product.
func (s *ProducstDefault) Save(p *internal.Product) (err error) {
	err = s.rp.Save(p)
	return
}