package repository

import (
	"database/sql"
	"errors"

	"github.com/usuario/repositorio/internal"

	"github.com/go-sql-driver/mysql"
)

// NewBuyerMysql creates a new instance of the buyer repository
func NewBuyerMysql(db *sql.DB) *BuyerMysql {
	return &BuyerMysql{db}
}

// BuyerMysql is the mysql implementation of the buyer repository
type BuyerMysql struct {
	// db is the database connection to mysql
	db *sql.DB
}

// FindAll returns all buyers from the database
func (r *BuyerMysql) FindAll() (buyers []internal.Buyer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `b.id`, `b.card_number_id`, `b.first_name`, `b.last_name` FROM `buyers` AS `b`")
	if err != nil {
		return
	}

	// iterate over the rows
	for rows.Next() {
		// create a new buyer
		var buyer internal.Buyer
		err = rows.Scan(&buyer.ID, &buyer.CardNumberID, &buyer.FirstName, &buyer.LastName)
		if err != nil {
			return
		}

		// append the buyer to the slice
		buyers = append(buyers, buyer)
	}

	// check for errors
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindByID returns a buyer from the database by its id
func (r *BuyerMysql) FindByID(id int) (buyer internal.Buyer, err error) {
	// execute the query
	row := r.db.QueryRow("SELECT `b.id`, `b.card_number_id`, `b.first_name`, `b.last_name` FROM `buyers` AS `b` WHERE `b.id` = ?", id)

	// scan the row into the buyer
	err = row.Scan(&buyer.ID, &buyer.CardNumberID, &buyer.FirstName, &buyer.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrBuyerRepositoryNotFound
		}
		return
	}

	return
}

// Save saves the given buyer in the database
func (r *BuyerMysql) Save(buyer *internal.Buyer) (err error) {
	// execute the query
	result, err := r.db.Exec(
		"INSERT INTO `buyers` (`card_number_id`, `first_name`, `last_name`) VALUES (?, ?, ?)",
		(*buyer).CardNumberID, (*buyer).FirstName, (*buyer).LastName,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrBuyerRepositoryDuplicated
			default:
				// ...
			}
			return
		}

		return
	}

	// get the id of the inserted buyer
	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	// set the id of the buyer
	(*buyer).ID = int(id)

	return
}

// Update updates the given buyer in the database
func (r *BuyerMysql) Update(buyer *internal.Buyer) (err error) {
	// execute the query
	_, err = r.db.Exec(
		"UPDATE `buyers` SET `card_number_id` = ?, `first_name` = ?, `last_name` = ? WHERE `id` = ?",
		(*buyer).CardNumberID, (*buyer).FirstName, (*buyer).LastName, (*buyer).ID,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrBuyerRepositoryDuplicated
			default:
				// ...
			}
			return
		}
	}

	return
}

// Delete deletes a buyer from the database by its id
func (r *BuyerMysql) Delete(id int) (err error) {
	// execute the query
	_, err = r.db.Exec("DELETE FROM `buyers` WHERE `id` = ?", id)
	if err != nil {
		return
	}

	return
}