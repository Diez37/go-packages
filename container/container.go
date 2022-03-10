package container

// Container general of container interface for any implementation
type Container interface {
	// Provide adding new object constructor
	Provide(interface{}) error

	// Provides adding new object constructors
	Provides(...interface{}) error

	// Invoke method call with dependency injection
	Invoke(interface{}) error
}
