package application

// Application is an interface for applications.
type Application interface {
	// Run runs the application.
	Run() (err error)
	// SetUp sets up the application.
	SetUp() (err error)
	// TearDown tears down the application.
	TearDown() (err error)
}