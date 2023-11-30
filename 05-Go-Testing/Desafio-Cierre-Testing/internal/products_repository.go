package internal

// RepositoryProducts is an interface that represents a repository.
type RepositoryProducts interface {
	// SearchProducts returns a list of products that match the query.
	SearchProducts(query ProductQuery) (p map[int]Product, err error)
}