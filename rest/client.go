package rest

import (
	"fmt"
	"net/url"
	"time"
)

type Client struct {
	Base url.URL

	key     string
	expires int64
}

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

func (c *Client) Auth(key string, expires int64) (*Client, error) {
	if expires < time.Now().Unix() {
		return nil, fmt.Errorf("expires expired in the past")
	}

	if key == "" {
		return nil, fmt.Errorf("key must not be empty")
	}

	c.key = key
	c.expires = expires

	return c, nil
}
