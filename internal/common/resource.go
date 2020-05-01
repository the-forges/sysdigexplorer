package common

// Resource provides a common face for interacting with an API Resource
type Resource interface {
	// Value takes a key, returning the value as an interface which should be
	// then type hinted in order to be used
	Value(string) interface{}
	// Set takes a key and value to be set for later use
	Set(string, interface{})
	// String returns the Resource details as a string
	String() string
}

// Searcher provides the basics to find one or more Resource
type Searcher interface {
	// Get returns a Resource or an error if it cannot be found
	Get() (Resource, error)
	// List returns a []Resource or an error if there is one
	List() ([]Resource, error)
}

// Writer provides the basics needed to create, update, or delete a Resource
type Writer interface {
	// Create writes a new Resource to the API
	// Update writes changes in a Resource to the API
	// Delete removes a Resource from the API
}
