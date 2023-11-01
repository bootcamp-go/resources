package internal

// ProductAttributes is an struct that represents the attributes of a product
type ProductAttributes struct {
	// Name is the name of the product
	Name string
	// Type is the type of the product
	Type string
	// Count is the quantity of the product
	Count int
	// Price is the price of the product
	Price float64
}

// Product is an struct that represents a product
type Product struct {
	// Id is the unique identifier of the product
	Id int
	// ProductAttributes represents the attributes of the product
	ProductAttributes
}