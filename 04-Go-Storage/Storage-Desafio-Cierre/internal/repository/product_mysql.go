package repository

import (
	"database/sql"
	"errors"

	"app/internal"
)

// NewRepositoryProductMySQL creates new mysql repository for product entity.
func NewRepositoryProductMySQL(db *sql.DB) *RepositoryProductMySQL {
	return &RepositoryProductMySQL{db}
}

// RepositoryProductMySQL is the MySQL repository implementation for product entity.
type RepositoryProductMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all products from the database.
func (r *RepositoryProductMySQL) FindAll() (p []internal.Product, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `description`, `price` FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var pr internal.Product
		// scan the row into the product
		err := rows.Scan(&pr.Id, &pr.Description, &pr.Price)
		if err != nil {
			return nil, err
		}
		// append the product to the slice
		p = append(p, pr)
	}

	return
}

// Save saves the product into the database.
func (r *RepositoryProductMySQL) Save(p *internal.Product) (err error) {
	// prepare the statement
	stmt, err := r.db.Prepare("INSERT INTO products (`description`, `price`) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute the statement
	res, err := stmt.Exec((*p).Description, (*p).Price)
	if err != nil {
		return err
	}

	// check the affected rows
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return errors.New("rows affected should be 1")
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*p).Id = int(id)

	return
}