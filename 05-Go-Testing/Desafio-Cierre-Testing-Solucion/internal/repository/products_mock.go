package repository

import "app/internal"

// NewProductsMock returns a new ProductsMock.
func NewProductsMock() *ProductsMock {
	return &ProductsMock{}
}

// ProductsMock is an struct that implements the Prosduct interface.
type ProductsMock struct {
	// FuncSearchProducts is the function that proxy the SearchProducts method.
	FuncSearchProducts func(query internal.ProductQuery) (p map[int]internal.Product, err error)
	// Spy
	Spy struct {
		// SearchProducts is the number of times the SearchProducts method is called.
		SearchProducts int
	}
}

// SearchProducts returns a list of products that match the query.
func (r *ProductsMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	// spy
	r.Spy.SearchProducts++

	// call the proxy function
	p, err = r.FuncSearchProducts(query)
	return
}