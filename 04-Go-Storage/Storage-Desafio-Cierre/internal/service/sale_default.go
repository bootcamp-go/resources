package service

import "app/internal"

// NewServiceSaleDefault creates new default service for sale entity.
func NewServiceSaleDefault(rp internal.RepositorySale) *ServiceSaleDefault {
	return &ServiceSaleDefault{rp}
}

// ServiceSaleDefault is the default service implementation for sale entity.
type ServiceSaleDefault struct {
	// rp is the repository for sale entity.
	rp internal.RepositorySale
}

// FindAll returns all sales.
func (sv *ServiceSaleDefault) FindAll() (s []internal.Sale, err error) {
	s, err = sv.rp.FindAll()
	return
}

// Save saves the sale.
func (sv *ServiceSaleDefault) Save(s *internal.Sale) (err error) {
	err = sv.rp.Save(s)
	return
}