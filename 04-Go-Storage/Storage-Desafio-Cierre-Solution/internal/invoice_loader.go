package internal

// LoaderInvoice is the interface that wraps the basic Load method.
type LoaderInvoice interface {
	// Load loads the invoice data from the source.
	Load() (i []Invoice, err error)
}