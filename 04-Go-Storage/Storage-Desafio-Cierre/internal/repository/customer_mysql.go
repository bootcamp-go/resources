package repository

import (
	"database/sql"
	"errors"

	"app/internal"
)

// NewRepositoryCustomerMySQL creates new mysql repository for customer entity.
func NewRepositoryCustomerMySQL(db *sql.DB) *RepositoryCustomerMySQL {
	return &RepositoryCustomerMySQL{db}
}

// RepositoryCustomerMySQL is the MySQL repository implementation for customer entity.
type RepositoryCustomerMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all customers from the database.
func (r *RepositoryCustomerMySQL) FindAll() (c []internal.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}

	return
}

// Save saves the customer into the database.
func (r *RepositoryCustomerMySQL) Save(c *internal.Customer) (err error) {
	// prepare the statement
	stmt, err := r.db.Prepare("INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute the statement
	res, err := stmt.Exec((*c).FirstName, (*c).LastName, (*c).Condition)
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
	(*c).Id = int(id)
	
	return
}

