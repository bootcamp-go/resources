package internal

import "time"

// ProductAttributes is a struct that contains the attributes of a product
type ProductAttributes struct {
	// Name is the name of the product
	Name string
	// Quantity is the quantity of the product
	Quantity int
	// CodeValue is the code value of the product
	CodeValue string
	// IsPublished is the published status of the product
	IsPublished bool
	// Expiration
	Expiration time.Time
	// Price
	Price float64
}

// Product is a struct that contains the attributes of a product
type Product struct {
	// Id is the unique identifier of the product
	Id int
	// ProductAttributes is the attributes of the product
	ProductAttributes
}