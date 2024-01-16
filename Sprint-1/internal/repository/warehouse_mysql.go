package repository

import (
	"database/sql"
	"errors"

	"github.com/usuario/repositorio/internal"

	"github.com/go-sql-driver/mysql"
)

// NewWarehouseMysql creates a new instance of the warehouse repository
func NewWarehouseMysql(db *sql.DB) *WarehouseMysql {
	return &WarehouseMysql{db}
}

// WarehouseMysql is the mysql implementation of the warehouse repository
type WarehouseMysql struct {
	// db is the database connection to mysql
	db *sql.DB
}

// FindAll returns all warehouses from the database
func (r *WarehouseMysql) FindAll() (warehouses []internal.Warehouse, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `w.id`, `w.warehouse_code`, `w.address`, `w.telephone`, `w.minimum_capacity`, `w.minimum_temperature` FROM `warehouses` AS `w`")
	if err != nil {
		return
	}

	// iterate over the rows
	for rows.Next() {
		// create a new warehouse
		var warehouse internal.Warehouse
		err = rows.Scan(&warehouse.ID, &warehouse.WarehouseCode, &warehouse.Address, &warehouse.Telephone, &warehouse.MinimumCapacity, &warehouse.MinimumTemperature)
		if err != nil {
			return
		}

		// append the warehouse to the slice
		warehouses = append(warehouses, warehouse)
	}

	// check for errors
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindByID returns a warehouse from the database by its id
func (r *WarehouseMysql) FindByID(id int) (warehouse internal.Warehouse, err error) {
	// execute the query
	row := r.db.QueryRow("SELECT `w.id`, `w.warehouse_code`, `w.address`, `w.telephone`, `w.minimum_capacity`, `w.minimum_temperature` FROM `warehouses` AS `w` WHERE `w.id` = ?", id)

	// scan the row into the warehouse
	err = row.Scan(&warehouse.ID, &warehouse.WarehouseCode, &warehouse.Address, &warehouse.Telephone, &warehouse.MinimumCapacity, &warehouse.MinimumTemperature)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrWarehouseRepositoryNotFound
		}
		return
	}

	return
}

// Save saves a warehouse into the database
func (r *WarehouseMysql) Save(warehouse *internal.Warehouse) (err error) {
	// execute the query
	result, err := r.db.Exec(
		"INSERT INTO `warehouses` (`warehouse_code`, `address`, `telephone`, `minimum_capacity`, `minimum_temperature`) VALUES (?, ?, ?, ?, ?)",
		(*warehouse).WarehouseCode, (*warehouse).Address, (*warehouse).Telephone, (*warehouse).MinimumCapacity, (*warehouse).MinimumTemperature,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrWarehouseRepositoryDuplicated
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

	// set the id of the warehouse
	(*warehouse).ID = int(id)

	return
}

// Update updates a warehouse in the database
func (r *WarehouseMysql) Update(warehouse *internal.Warehouse) (err error) {
	// execute the query
	_, err = r.db.Exec(
		"UPDATE `warehouses` AS `w` SET `w.warehouse_code` = ?, `w.address` = ?, `w.telephone` = ?, `w.minimum_capacity` = ?, `w.minimum_temperature` = ? WHERE `w.id` = ?",
		(*warehouse).WarehouseCode, (*warehouse).Address, (*warehouse).Telephone, (*warehouse).MinimumCapacity, (*warehouse).MinimumTemperature, (*warehouse).ID,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrWarehouseRepositoryDuplicated
			default:
				// ...
			}
			return
		}

		return
	}

	return
}

// Delete deletes a warehouse from the database
func (r *WarehouseMysql) Delete(id int) (err error) {
	// execute the query
	_, err = r.db.Exec("DELETE FROM `warehouses` WHERE `id` = ?", id)
	if err != nil {
		return
	}

	return
}