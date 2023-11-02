package internal

// UserAttributes is an struct that represents the attributes of a user
type UserAttributes struct {
	// Name is the name of the user
	Name string
	// Age is the age of the user
	Age int
	// Email is the email of the user
	Email string
}

// User is an struct that represents a user
type User struct {
	// Id is the unique identifier of the user
	Id int
	// UserAttributes is the attributes of the user
	UserAttributes
}