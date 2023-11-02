package internal

import "errors"

var (
	// ErrFieldDuplicated is an error that represent a duplicated field.
	ErrFieldDuplicated = errors.New("field duplicated")
)

// UserRepository is an interface that represents a repository of users
type UserRepository interface {
	// FindAll returns all the users
	FindAll() (u map[int]User, err error)

	// Save saves a user
	Save(u *User) (err error)
}