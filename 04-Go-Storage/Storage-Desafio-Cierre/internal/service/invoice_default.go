package service

import "app/internal"

// NewServiceInvoiceDefault creates new default service for invoice entity.
func NewServiceInvoiceDefault(rp internal.RepositoryInvoice) *ServiceInvoiceDefault {
	return &ServiceInvoiceDefault{rp}
}

// ServiceInvoiceDefault is the default service implementation for invoice entity.
type ServiceInvoiceDefault struct {
	// rp is the repository for invoice entity.
	rp internal.RepositoryInvoice
}

// FindAll returns all invoices.
func (s *ServiceInvoiceDefault) FindAll() (i []internal.Invoice, err error) {
	i, err = s.rp.FindAll()
	return
}

// Save saves the invoice.
func (s *ServiceInvoiceDefault) Save(i *internal.Invoice) (err error) {
	err = s.rp.Save(i)
	return
}