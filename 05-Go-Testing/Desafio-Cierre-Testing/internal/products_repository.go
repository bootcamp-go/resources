package internal

// RepositoryProduct is an interface that represents a repository.
type RepositoryProduct interface {
	// SearchProducts returns a list of products that match the query.
	SearchProducts(query ProductQuery) (p map[int]Product, err error)
}