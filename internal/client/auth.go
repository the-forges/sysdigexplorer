package client

import (
	"net/url"

	"github.ibm.com/ZaaS/sysdigexplorer/internal/errors"
)

// Client encapsulates connection details
type Client struct {
	url    string
	apiKey string
}

// NewClient will give you a *Client with all the required info to make client
// calls to a given sysdig instance
func NewClient(url, apiKey string) (*Client, error) {
	if url == "" {
		return nil, errors.New(ErrMsgMissingURL)
	}
	if apiKey == "" {
		return nil, errors.New(ErrMsgMissingAPIKey)
	}
	return &Client{url: url, apiKey: apiKey}, nil
}

// URL returns a URL; unless there was a problem parsing it, in which case
// it returns an error
func (c *Client) URL() (*url.URL, error) {
	return url.Parse(c.url)
}

// APIKey returns the API Key you are using, or an error if it is missing
func (c *Client) APIKey() (string, error) {
	if c.apiKey == "" {
		return "", errors.New(ErrMsgMissingAPIKey)
	}
	return c.apiKey, nil
}
