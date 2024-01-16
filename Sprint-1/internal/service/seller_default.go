package service

import "github.com/usuario/repositorio/internal"

// NewSellerDefault creates a new instance of the seller service
func NewSellerDefault(rp internal.SellerRepository) *SellerDefault {
	return &SellerDefault{
		rp: rp,
	}
}

// SellerDefault is the default implementation of the seller service
type SellerDefault struct {
	// rp is the repository used by the service
	rp internal.SellerRepository
}

// FindAll returns all sellers
func (s *SellerDefault) FindAll() (sellers []internal.Seller, err error) {
	return
}

// FindByID returns a seller
func (s *SellerDefault) FindByID(id int) (seller internal.Seller, err error) {
	return
}

// Save creates a new seller
func (s *SellerDefault) Save(seller *internal.Seller) (err error) {
	return
}

// Update updates a seller
func (s *SellerDefault) Update(seller *internal.Seller) (err error) {
	return
}

// Delete deletes a seller
func (s *SellerDefault) Delete(id int) (err error) {
	return
}