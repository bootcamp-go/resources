package internal

// LoaderCustomer is the interface that wraps the basic Load method.
type LoaderCustomer interface {
	// Load loads the customer data from the source.
	Load() (c []Customer, err error)
}