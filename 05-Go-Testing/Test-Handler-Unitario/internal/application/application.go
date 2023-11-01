package application

// Application in an interface that represents an application
type Application interface {
	// Run runs an application
	Run() (err error)
	// SetUp sets up an application
	SetUp() (err error)
	// TearDown tears down an application
	TearDown() (err error)
}