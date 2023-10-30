package repository

import (
	"database/sql"
	"errors"

	"app/internal"
)

// NewRepositoryInvoiceMySQL creates new mysql repository for invoice entity.
func NewRepositoryInvoiceMySQL(db *sql.DB) *RepositoryInvoiceMySQL {
	return &RepositoryInvoiceMySQL{db}
}

// RepositoryInvoiceMySQL is the MySQL repository implementation for invoice entity.
type RepositoryInvoiceMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all invoices from the database.
func (r *RepositoryInvoiceMySQL) FindAll() (i []internal.Invoice, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `datetime`, `total`, `customer_id` FROM invoices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var iv internal.Invoice
		// scan the row into the invoice
		err := rows.Scan(&iv.Id, &iv.Datetime, &iv.Total, &iv.CustomerId)
		if err != nil {
			return nil, err
		}
		// append the invoice to the slice
		i = append(i, iv)
	}

	return
}

// Save saves the invoice into the database.
func (r *RepositoryInvoiceMySQL) Save(i *internal.Invoice) (err error) {
	// prepare the statement
	stmt, err := r.db.Prepare("INSERT INTO invoices (`datetime`, `total`, `customer_id`) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute the statement
	res, err := stmt.Exec((*i).Datetime, (*i).Total, (*i).CustomerId)
	if err != nil {
		return err
	}

	// check the affected rows
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("no rows affected")
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*i).Id = int(id)

	return
}