package internal

// LoaderSale is the interface that wraps the basic Load method.
type LoaderSale interface {
	// Load loads the sale data from the source.
	Load() (s []Sale, err error)
}