package repository

import "app/internal"

// NewRepositoryProductMock returns a new RepositoryProductMock.
func NewRepositoryProductMock() *RepositoryProductMock {
	return &RepositoryProductMock{}
}

// RepositoryProductMock is an struct that implements the RepositoryProduct interface.
type RepositoryProductMock struct {
	// FuncSearchProducts is the function that proxy the SearchProducts method.
	FuncSearchProducts func(query internal.ProductQuery) (p map[int]internal.Product, err error)
	// Spy
	Spy struct {
		// SearchProducts is the number of times the SearchProducts method is called.
		SearchProducts int
	}
}

// SearchProducts returns a list of products that match the query.
func (r *RepositoryProductMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	// spy
	r.Spy.SearchProducts++

	// call the proxy function
	p, err = r.FuncSearchProducts(query)
	return
}