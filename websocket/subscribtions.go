package websocket

import "github.com/pkg/errors"

func (c *Client) Subscribe(name, symbol string) error {
	var m = WSEvent{
		Op: "subscribe",
		Args: []interface{}{
			name + ":" + symbol,
		},
	}

	if err := c.SendJSON(m); err != nil {
		return errors.Wrap(err, "can't send json to server")
	}

	return nil
}

func (c *Client) Unsubscribe(name, symbol string) error {
	var m = WSEvent{
		Op: "unsubscribe",
		Args: []interface{}{
			name + ":" + symbol,
		},
	}

	if err := c.SendJSON(m); err != nil {
		return errors.Wrap(err, "can't send json to server")
	}

	return nil
}
