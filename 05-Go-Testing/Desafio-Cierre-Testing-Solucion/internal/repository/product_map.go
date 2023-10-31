package repository

import "app/internal"

// NewRepositoryProductMap returns a new RepositoryProductMap.
func NewRepositoryProductMap(db map[int]internal.Product) *RepositoryProductMap {
	// default values
	defaultDb := make(map[int]internal.Product)
	if db != nil {
		defaultDb = db
	}

	return &RepositoryProductMap{
		db: defaultDb,
	}
}

// ProductAttributes is an struct that implements the RepositoryProduct interface.
type RepositoryProductMap struct {
	// db is the map of products.
	db map[int]internal.Product
}

// SearchProducts returns a list of products that match the query.
func (r *RepositoryProductMap) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	p = make(map[int]internal.Product)

	// search the products
	for k, v := range r.db {
		// check if each query field is set
		if query.Id > 0 && query.Id != k {
			continue
		}

		// add the product to the result
		p[k] = v
	}

	return
}