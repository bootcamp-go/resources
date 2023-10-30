package migrator

import (
	"app/internal"
)

// NewMigratorInvoiceDatabase returns a new MigratorInvoiceToDatabase
func NewMigratorInvoiceToDatabase(ld internal.LoaderInvoice, rp internal.RepositoryInvoice) (m *MigratorInvoiceToDatabase) {
	m = &MigratorInvoiceToDatabase{
		ld: ld,
		rp: rp,
	}
	return
}

// MigratorInvoiceToDatabase is the implementation of the interface MigratorInvoice
type MigratorInvoiceToDatabase struct {
	// ld is the loader to load the data
	ld internal.LoaderInvoice
	// rp is the repository to access the database
	rp internal.RepositoryInvoice
}

// Migrate migrates the data from the a source to a destination
func (m *MigratorInvoiceToDatabase) Migrate() (err error) {
	// load the data
	i, err := m.ld.Load()
	if err != nil {
		return
	}

	// save each customer
	for _, v := range i {
		err = m.rp.Save(&v)
		if err != nil {
			return
		}
	}

	return
}