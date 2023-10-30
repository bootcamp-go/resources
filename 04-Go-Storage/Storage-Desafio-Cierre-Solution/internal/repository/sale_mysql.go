package repository

import (
	"database/sql"
	"errors"

	"app/internal"
)

// NewRepositorySaleMySQL creates new mysql repository for sale entity.
func NewRepositorySaleMySQL(db *sql.DB) *RepositorySaleMySQL {
	return &RepositorySaleMySQL{db}
}

// RepositorySaleMySQL is the MySQL repository implementation for sale entity.
type RepositorySaleMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all sales from the database.
func (r *RepositorySaleMySQL) FindAll() (s []internal.Sale, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `quantity`, `product_id`, `invoice_id` FROM sales")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var sa internal.Sale
		// scan the row into the sale
		err := rows.Scan(&sa.Id, &sa.Quantity, &sa.ProductId, &sa.InvoiceId)
		if err != nil {
			return nil, err
		}
		// append the sale to the slice
		s = append(s, sa)
	}

	return
}

// Save saves the sale into the database.
func (r *RepositorySaleMySQL) Save(s *internal.Sale) (err error) {
	// prepare the statement
	stmt, err := r.db.Prepare("INSERT INTO sales (`quantity`, `product_id`, `invoice_id`) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute the statement
	res, err := stmt.Exec((*s).Quantity, (*s).ProductId, (*s).InvoiceId)
	if err != nil {
		return err
	}

	// check the affected rows
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("no rows affected")
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*s).Id = int(id)

	return
}