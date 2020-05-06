package common

import (
	"net/url"
)

// Client ecapsulates client details
type Client interface {
	// URL returns a URL; unless there was a problem parsing it, in which case
	// it returns an error
	URL() (*url.URL, error)
	// APIKey returns the API Key you are using, or an error if it is missing
	APIKey() (string, error)
	// Valid ensures URL() and APIKey return no errors
	Valid() bool
	// Request takes a method, an endpoint, and a map of additional options in
	// order to use the API
	Request(string, string, map[string]interface{}) ([]byte, error)
}
