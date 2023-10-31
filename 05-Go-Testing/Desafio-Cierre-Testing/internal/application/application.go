package application

// Application is an interface that represents the application.
type Application interface {
	// Run runs the application.
	Run() (err error)
	// SetUp sets up the application.
	SetUp() (err error)
	// TearDown tears down the application.
	// - should be used as a defer function
	TearDown() (err error)
}