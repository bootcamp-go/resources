package service

import "github.com/usuario/repositorio/internal"

// NewBuyerDefault creates a new instance of the buyer service
func NewBuyerDefault(rp internal.BuyerRepository) *BuyerDefault {
	return &BuyerDefault{
		rp: rp,
	}
}

// BuyerDefault is the default implementation of the buyer service
type BuyerDefault struct {
	// rp is the repository used by the service
	rp internal.BuyerRepository
}

// FindAll returns all buyers
func (s *BuyerDefault) FindAll() (buyers []internal.Buyer, err error) {
	return
}

// FindByID returns a buyer
func (s *BuyerDefault) FindByID(id int) (buyer internal.Buyer, err error) {
	return
}

// Save creates a new buyer
func (s *BuyerDefault) Save(buyer *internal.Buyer) (err error) {
	return
}

// Update updates a buyer
func (s *BuyerDefault) Update(buyer *internal.Buyer) (err error) {
	return
}

// Delete deletes a buyer
func (s *BuyerDefault) Delete(id int) (err error) {
	return
}