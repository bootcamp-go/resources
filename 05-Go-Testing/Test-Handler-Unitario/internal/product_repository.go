package internal

import "errors"

var (
	// ErrProductDuplicatedField is an error that is returned when a product has a duplicated field
	ErrProductDuplicatedField = errors.New("product field is duplicated")
)

// ProductRepository is an interface that represents a product repository
type ProductRepository interface {
	// Save saves a product
	Save(product *Product) (err error)
}