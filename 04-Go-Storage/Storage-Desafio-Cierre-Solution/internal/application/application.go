package application

// Application is an interface that represents an application.
type Application interface {
	// TearDown tears down the application.
	TearDown()
	// SetUp sets up the application.
	SetUp() (err error)
	// Run runs the application.
	Run() (err error)
}