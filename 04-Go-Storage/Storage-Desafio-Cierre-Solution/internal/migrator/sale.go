package migrator

import (
	"app/internal"
)

// NewMigratorSaleDatabase returns a new MigratorSaleToDatabase
func NewMigratorSaleToDatabase(ld internal.LoaderSale, rp internal.RepositorySale) (m *MigratorSaleToDatabase) {
	m = &MigratorSaleToDatabase{
		ld: ld,
		rp: rp,
	}
	return
}

// MigratorSaleToDatabase is the implementation of the interface MigratorSale
type MigratorSaleToDatabase struct {
	// ld is the loader to load the data
	ld internal.LoaderSale
	// rp is the repository to access the database
	rp internal.RepositorySale
}

// Migrate migrates the data from the a source to a destination
func (m *MigratorSaleToDatabase) Migrate() (err error) {
	// load the data
	s, err := m.ld.Load()
	if err != nil {
		return
	}

	// save each customer
	for _, v := range s {
		err = m.rp.Save(&v)
		if err != nil {
			return
		}
	}

	return
}