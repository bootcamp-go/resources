package application

// Application is an interface that represents an application
type Application interface {
	// SetUp is a method that sets up the application
	SetUp() (err error)
	// Run is a method that runs the application
	Run() (err error)
}