package client

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/the-forges/sysdigexplorer/errors"
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
	if !c.Valid() {
		return "", errors.New(ErrMsgInvalidClient)
	}
	return c.apiKey, nil
}

// Valid ensures URL() and APIKey return no errors
func (c *Client) Valid() bool {
	if c.apiKey == "" {
		return false
	}
	if c.url == "" {
		return false
	}
	return true
}

// Request takes a method, an endpoint, and a map of additional options in
// order to use the API
func (c *Client) Request(method, endpoint string, options map[string]interface{}) ([]byte, error) {
	url := c.url + endpoint
	r, err := http.NewRequest(method, url, nil) // TODO: request body
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+c.apiKey)

	client := http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
