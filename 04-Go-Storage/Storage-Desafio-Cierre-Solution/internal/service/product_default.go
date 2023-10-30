package service

import "app/internal"

// NewServiceProductDefault creates new default service for product entity.
func NewServiceProductDefault(rp internal.RepositoryProduct) *ServiceProductDefault {
	return &ServiceProductDefault{rp}
}

// ServiceProductDefault is the default service implementation for product entity.
type ServiceProductDefault struct {
	// rp is the repository for product entity.
	rp internal.RepositoryProduct
}

// FindAll returns all products.
func (s *ServiceProductDefault) FindAll() (p []internal.Product, err error) {
	p, err = s.rp.FindAll()
	return
}

// FindTopProductsByAmountSold returns the top products by amount sold.
func (s *ServiceProductDefault) FindTopProductsByAmountSold(limit int) (p []internal.ProductAmountSold, err error) {
	p, err = s.rp.FindTopProductsByAmountSold(limit)
	return
}

// Save saves the product.
func (s *ServiceProductDefault) Save(p *internal.Product) (err error) {
	err = s.rp.Save(p)
	return
}