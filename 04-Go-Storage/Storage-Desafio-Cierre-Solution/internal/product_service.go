package internal

// ServiceProduct is the interface that wraps the basic Product methods.
type ServiceProduct interface {
	// FindAll returns all products.
	FindAll() (p []Product, err error)
	// FindTopProductsByAmountSold returns the top products by amount sold.
	FindTopProductsByAmountSold(limit int) (p []ProductAmountSold, err error)
	// Save saves a product.
	Save(p *Product) (err error)
}