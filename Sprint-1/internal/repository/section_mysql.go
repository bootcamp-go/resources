package repository

import (
	"database/sql"
	"errors"

	"github.com/usuario/repositorio/internal"

	"github.com/go-sql-driver/mysql"
)

// NewSectionMysql creates a new instance of the section repository
func NewSectionMysql(db *sql.DB) *SectionMysql {
	return &SectionMysql{db}
}

// SectionMysql is the mysql implementation of the section repository
type SectionMysql struct {
	// db is the database connection to mysql
	db *sql.DB
}

// FindAll returns all sections from the database
func (r *SectionMysql) FindAll() (sections []internal.Section, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `s.id`, `s.section_number`, `s.current_temperature`, `s.minimum_temperature`, `s.current_capacity`, `s.minimum_capacity`, `s.maximum_capacity`, `s.warehouse_id`, `s.product_type_id` FROM `sections` AS `s`")
	if err != nil {
		return
	}

	// iterate over the rows
	for rows.Next() {
		// create a new section
		var section internal.Section
		err = rows.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
		if err != nil {
			return
		}

		// append the section to the slice
		sections = append(sections, section)
	}

	// check for errors
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindByID returns a section from the database by its id
func (r *SectionMysql) FindByID(id int) (section internal.Section, err error) {
	// execute the query
	row := r.db.QueryRow("SELECT `s.id`, `s.section_number`, `s.current_temperature`, `s.minimum_temperature`, `s.current_capacity`, `s.minimum_capacity`, `s.maximum_capacity`, `s.warehouse_id`, `s.product_type_id` FROM `sections` AS `s` WHERE `s.id` = ?", id)

	// scan the row into the section
	err = row.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrSectionRepositoryNotFound
		}
		return
	}

	return
}

// Save saves a section into the database
func (r *SectionMysql) Save(section *internal.Section) (err error) {
	// execute the query
	result, err := r.db.Exec(
		"INSERT INTO `sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		(*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrSectionRepositoryDuplicated
			default:
				// ...
			}
			return
		}

		return
	}

	// get the id of the inserted section
	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	// set the id of the section
	(*section).ID = int(id)

	return
}

// Update updates a section in the database
func (r *SectionMysql) Update(section *internal.Section) (err error) {
	// execute the query
	_, err = r.db.Exec(
		"UPDATE `sections` SET `section_number` = ?, `current_temperature` = ?, `minimum_temperature` = ?, `current_capacity` = ?, `minimum_capacity` = ?, `maximum_capacity` = ?, `warehouse_id` = ?, `product_type_id` = ? WHERE `id` = ?",
		(*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID, (*section).ID,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrSectionRepositoryDuplicated
			default:
				// ...
			}
			return
		}

		return
	}

	return
}

// Delete deletes a section from the database by its id
func (r *SectionMysql) Delete(id int) (err error) {
	// execute the query
	_, err = r.db.Exec("DELETE FROM `sections` WHERE `id` = ?", id)
	if err != nil {
		return
	}

	return
}