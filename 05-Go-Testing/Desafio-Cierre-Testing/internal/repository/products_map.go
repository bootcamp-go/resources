package repository

import "app/internal"

// NewRepositoryProductsMap returns a new RepositoryProductsMap.
func NewRepositoryProductsMap(db map[int]internal.Product) *RepositoryProductsMap {
	// default values
	defaultDb := make(map[int]internal.Product)
	if db != nil {
		defaultDb = db
	}

	return &RepositoryProductsMap{
		db: defaultDb,
	}
}

// ProductAttributes is an struct that implements the RepositoryProduct interface.
type RepositoryProductsMap struct {
	// db is the map of products.
	db map[int]internal.Product
}

// SearchProducts returns a list of products that match the query.
func (r *RepositoryProductsMap) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	p = make(map[int]internal.Product)

	// search the products
	for k, v := range r.db {
		// check if each query field is set
		if query.Id > 0 && query.Id != v.Id {
			continue
		}

		// add the product to the result
		p[k] = v
	}

	return
}