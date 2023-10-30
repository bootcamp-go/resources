package internal

// Migrator is the interface that wraps the basic Migrate method
type Migrator interface {
	// Migrate migrates the data from the a source to a destination
	Migrate() (err error)
}