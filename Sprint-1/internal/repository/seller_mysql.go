package repository

import (
	"database/sql"
	"errors"

	"github.com/usuario/repositorio/internal"

	"github.com/go-sql-driver/mysql"
)

// NewSellerMysql creates a new instance of the seller repository
func NewSellerMysql(db *sql.DB) *SellerMysql {
	return &SellerMysql{db}
}

// SellerMysql is the mysql implementation of the seller repository
type SellerMysql struct {
	// db is the database connection to mysql
	db *sql.DB
}

// FindAll returns all sellers from the database
func (r *SellerMysql) FindAll() (sellers []internal.Seller, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `s.id`, `s.cid`, `c.company_name`, `c.address`, `c.telephone` FROM `sellers` AS `s`")
	if err != nil {
		return
	}

	// iterate over the rows
	for rows.Next() {
		// create a new seller
		var seller internal.Seller
		err = rows.Scan(&seller.ID, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone)
		if err != nil {
			return
		}

		// append the seller to the slice
		sellers = append(sellers, seller)
	}

	// check for errors
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindByID returns a seller from the database by its id
func (r *SellerMysql) FindByID(id int) (seller internal.Seller, err error) {
	// execute the query
	row := r.db.QueryRow("SELECT `s.id`, `s.cid`, `c.company_name`, `c.address`, `c.telephone` FROM `sellers` AS `s` WHERE `s.id` = ?", id)

	// scan the row into the seller
	err = row.Scan(&seller.ID, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrSellerRepositoryNotFound
			return
		}
		return
	}

	return
}

// Save saves a seller into the database
func (r *SellerMysql) Save(seller *internal.Seller) (err error) {
	// execute the query
	result, err := r.db.Exec(
		"INSERT INTO `sellers` (`cid`, `company_name`, `address`, `telephone`) VALUES (?, ?, ?, ?)",
		(*seller).CID, (*seller).CompanyName, (*seller).Address, (*seller).Telephone,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrSellerRepositoryDuplicated
			default:
				// ...
			}
			return
		}

		return
	}

	// get the last inserted id
	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	// set the id of the seller
	(*seller).ID = int(id)

	return
}

// Update updates a seller in the database
func (r *SellerMysql) Update(seller *internal.Seller) (err error) {
	// execute the query
	_, err = r.db.Exec(
		"UPDATE `sellers` SET `cid` = ?, `company_name` = ?, `address` = ?, `telephone` = ? WHERE `id` = ?",
		(*seller).CID, (*seller).CompanyName, (*seller).Address, (*seller).Telephone, (*seller).ID,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrSellerRepositoryDuplicated
			default:
				// ...
			}
			return
		}

		return
	}

	return
}

// Delete deletes a seller from the database
func (r *SellerMysql) Delete(id int) (err error) {
	// execute the query
	_, err = r.db.Exec("DELETE FROM `sellers` WHERE `id` = ?", id)
	if err != nil {
		return
	}

	return
}