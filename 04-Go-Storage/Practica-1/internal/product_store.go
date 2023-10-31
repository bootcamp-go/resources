package internal

// StoreProduct is an interface for a product store.
type StoreProduct interface {
	// ReadAll reads all products from the store.
	ReadAll() (p map[int]Product, err error)
	// WriteAll writes all products to the store.
	WriteAll(p map[int]Product) (err error)
}