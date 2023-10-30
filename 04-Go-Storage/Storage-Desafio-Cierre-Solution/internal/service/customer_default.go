package service

import "app/internal"

// NewServiceCustomerDefault creates new default service for customer entity.
func NewServiceCustomerDefault(rp internal.RepositoryCustomer) *ServiceCustomerDefault {
	return &ServiceCustomerDefault{rp}
}

// ServiceCustomerDefault is the default service implementation for customer entity.
type ServiceCustomerDefault struct {
	// rp is the repository for customer entity.
	rp internal.RepositoryCustomer
}

// FindAll returns all customers.
func (s *ServiceCustomerDefault) FindAll() (c []internal.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

// FindTopActiveCustomersByAmountSpent returns the top active customers by amount spent.
func (s *ServiceCustomerDefault) FindTopActiveCustomersByAmountSpent(limit int) (c []internal.CustomerSpent, err error) {
	c, err = s.rp.FindTopActiveCustomersByAmountSpent(limit)
	return
}

// FindInvoicesByCondition returns the total invoices by customer condition.
func (s *ServiceCustomerDefault) FindInvoicesByCondition() (c []internal.CustomerInvoicesByCondition, err error) {
	c, err = s.rp.FindInvoicesByCondition()
	return
}

// Save saves the customer.
func (s *ServiceCustomerDefault) Save(c *internal.Customer) (err error) {
	err = s.rp.Save(c)
	return
}