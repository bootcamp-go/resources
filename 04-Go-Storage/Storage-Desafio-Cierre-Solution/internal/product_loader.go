package internal

// LoaderProduct is the interface that wraps the basic Load method.
type LoaderProduct interface {
	// Load loads the product data from the source.
	Load() (p []Product, err error)
}