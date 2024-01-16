package repository

import (
	"database/sql"
	"errors"

	"github.com/usuario/repositorio/internal"

	"github.com/go-sql-driver/mysql"
)

// NewEmployeeMysql creates a new instance of the employee repository
func NewEmployeeMysql(db *sql.DB) *EmployeeMysql {
	return &EmployeeMysql{db}
}

// EmployeeMysql is the mysql implementation of the employee repository
type EmployeeMysql struct {
	// db is the database connection to mysql
	db *sql.DB
}

// FindAll returns all employees from the database
func (r *EmployeeMysql) FindAll() (employees []internal.Employee, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `e.id`, `e.card_number_id`, `e.first_name`, `e.last_name`, `e.warehouse_id` FROM `employees` AS `e`")
	if err != nil {
		return
	}

	// iterate over the rows
	for rows.Next() {
		// create a new employee
		var employee internal.Employee
		err = rows.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID)
		if err != nil {
			return
		}

		// append the employee to the slice
		employees = append(employees, employee)
	}

	// check for errors
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindByID returns a employee from the database by its id
func (r *EmployeeMysql) FindByID(id int) (employee internal.Employee, err error) {
	// execute the query
	row := r.db.QueryRow("SELECT `e.id`, `e.card_number_id`, `e.first_name`, `e.last_name`, `e.warehouse_id` FROM `employees` AS `e` WHERE `e.id` = ?", id)

	// scan the row into the employee
	err = row.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrEmployeeRepositoryNotFound
		}
		return
	}

	return
}

// Save saves the given employee in the database
func (r *EmployeeMysql) Save(employee *internal.Employee) (err error) {
	// execute the query
	result, err := r.db.Exec(
		"INSERT INTO `employees` (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES (?, ?, ?, ?)",
		(*employee).CardNumberID, (*employee).FirstName, (*employee).LastName, (*employee).WarehouseID,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrEmployeeRepositoryDuplicated
			default:
				// ...
			}
			return
		}

		return
	}

	// get the id of the inserted employee
	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	// set the id of the employee
	(*employee).ID = int(id)

	return
}

// Update updates the given employee in the database
func (r *EmployeeMysql) Update(employee *internal.Employee) (err error) {
	// execute the query
	_, err = r.db.Exec(
		"UPDATE `employees` SET `card_number_id` = ?, `first_name` = ?, `last_name` = ?, `warehouse_id` = ? WHERE `id` = ?",
		(*employee).CardNumberID, (*employee).FirstName, (*employee).LastName, (*employee).WarehouseID, (*employee).ID,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				err = internal.ErrEmployeeRepositoryDuplicated
			default:
				// ...
			}
			return
		}

		return
	}

	return
}

// Delete deletes the given employee from the database
func (r *EmployeeMysql) Delete(id int) (err error) {
	// execute the query
	_, err = r.db.Exec("DELETE FROM `employees` WHERE `id` = ?", id)
	if err != nil {
		return
	}

	return
}