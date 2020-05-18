package rest

import (
	"fmt"
	"net/url"
	"time"
)

// Client stores the requiared data for requesting
type Client struct {
	Base url.URL

	key     string
	expires int64
}

// NewClient creates new client for requesting data from the bitmex exchange
func NewClient(test bool) *Client {
	host := BitmexHost
	if test {
		host = BitmexHostTestnet
	}

	return &Client{
		Base: url.URL{
			Scheme: "https",
			Host:   host,
			Path:   BitmexAPIPath,
		},
	}
}

// Auth return authoarized client
// it takes public key and expires time in unix format
// the secret key you need to pass with requesting as a context value
func (c *Client) Auth(key string, expires int64) (*Client, error) {
	if expires < time.Now().Unix() {
		return nil, fmt.Errorf("expired in the past")
	}

	if key == "" {
		return nil, fmt.Errorf("key should not be empty")
	}

	c.key = key
	c.expires = expires

	return c, nil
}
