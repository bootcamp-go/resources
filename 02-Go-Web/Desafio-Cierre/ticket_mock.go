package repository

import "context"

// NewRepositoryTicketMock creates a new repository for tickets in a map
func NewRepositoryTicketMock() *RepositoryTicketMap {
	return &RepositoryTicketMap{}
}

// RepositoryTicketMock implements the repository interface for tickets
type RepositoryTicketMock struct {
	// FuncGet represents the mock for the Get function
	FuncGet func() (t map[int]internal.TicketAttributes, err error)
	// FuncGetTicketsByDestinationCountry
	FuncGetTicketsByDestinationCountry func(country string) (t map[int]internal.TicketAttributes, err error)

	// Spy verifies if the methods were called
	Spy struct {
		// Get represents the spy for the Get function
		Get int
		// GetTicketsByDestinationCountry represents the spy for the GetTicketsByDestinationCountry function
		GetTicketsByDestinationCountry int
	}
}

// GetAll returns all the tickets
func (r *RepositoryTicketMock) Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error) {
	// spy
	r.Spy.Get++

	// mock
	t, err = r.FuncGet()
	return
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMock) GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error) {
	// spy
	r.Spy.GetTicketsByDestinationCountry++

	// mock
	t, err = r.FuncGetTicketsByDestinationCountry(country)
	return
}