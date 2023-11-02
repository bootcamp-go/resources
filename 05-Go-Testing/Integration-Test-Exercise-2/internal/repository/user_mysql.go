package repository

import (
	"app/internal"
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// UserMySQL is the MySQL implementation of the user repository.
type UserMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// NewUserMySQL returns a new instance of UserMySQL.
func NewUserMySQL(db *sql.DB) *UserMySQL {
	return &UserMySQL{db}
}

// FindAll returns all the users.
func (r *UserMySQL) FindAll() (u map[int]internal.User, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `name`, `age`, `email` FROM `users`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	u = make(map[int]internal.User)
	for rows.Next() {
		var us internal.User
		err = rows.Scan(&us.Id, &us.Name, &us.Age, &us.Email)
		if err != nil {
			return nil, err
		}

		u[us.Id] = us
	}

	return
}

// Save saves a user.
func (r *UserMySQL) Save(u *internal.User) (err error) {
	// prepare the statement
	stmt, err := r.db.Prepare("INSERT INTO `users` (`id`, `name`, `age`, `email`) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute the statement
	res, err := stmt.Exec(u.Id, u.Name, u.Age, u.Email)
	if err != nil {
		var errMySQL *mysql.MySQLError
		if errors.As(err, &errMySQL) {
			switch errMySQL.Number {
			case 1062:
				err = internal.ErrFieldDuplicated
			}
			return
		}
		return
	}

	// check the number of rows affected
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("rows affected should be 1")
	}

	return
}