package repository

import "app/internal"

// NewProductMock return a new ProductMock
func NewProductMock() *ProductMock {
	return &ProductMock{
		FuncSave: func(product *internal.Product) (err error) {
			return
		},
	}
}

// ProductMock is a mock for product repository
type ProductMock struct {
	// FuncSave proxy the Save method
	FuncSave func(product *internal.Product) (err error)
	// Calls count the number of calls of each method
	Calls struct {
		// Save count the number of calls of Save method
		Save int
	}
}

// Save mock the Save method
func (m *ProductMock) Save(product *internal.Product) (err error) {
	// increment the calls count
	m.Calls.Save++

	// call the proxy method
	err = m.FuncSave(product)
	return
}