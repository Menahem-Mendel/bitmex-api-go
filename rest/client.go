package rest

import (
	"fmt"
	"net/url"

	bitmex "github.com/Menahem-Mendel/bitmex-api-go"
)

// Client stores the requiared data for requesting
type Client struct {
	Base url.URL

	key string
}

// NewClient creates new client for requesting data from the bitmex exchange
func NewClient(test bool) *Client {
	host := bitmex.BitmexHost
	if test {
		host = bitmex.BitmexHostTestnet
	}

	return &Client{
		Base: url.URL{
			Scheme: "https",
			Host:   host,
			Path:   bitmex.BitmexAPIPathV1,
		},
	}
}

// NewAuthClient return authoarized client
// it takes public key and expires time in unix format
// the secret key you need to pass with requesting as a context value
func NewAuthClient(test bool, key string) (*Client, error) {
	if key == "" {
		return nil, fmt.Errorf("key should not be empty")
	}
	host := bitmex.BitmexHost
	if test {
		host = bitmex.BitmexHostTestnet
	}

	return &Client{
		Base: url.URL{
			Scheme: "https",
			Host:   host,
			Path:   bitmex.BitmexAPIPathV1,
		},
		key: key,
	}, nil
}
