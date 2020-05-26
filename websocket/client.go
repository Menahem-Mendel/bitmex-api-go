package websocket

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type Client struct {
	mu sync.Mutex

	rc chan []byte
	wc chan []byte

	key    string
	secret string

	*websocket.Conn
	url.URL
}

func New(host, path string) *Client {
	c := Client{
		URL: url.URL{
			Scheme: "wss",
			Host:   host,
			Path:   path,
		},

		rc: make(chan []byte),
		wc: make(chan []byte),
	}

	return &c
}

func (c *Client) Auth(key, secret string) *Client {
	c.key = key
	c.secret = secret

	return c
}

func (c *Client) Connect() error {
	header := make(http.Header)

	header.Set("content-type", "application/json")
	header.Set("accept", "application/json")
	header.Set("user-agent", "Jpro")

	conn, resp, err := websocket.DefaultDialer.Dial(c.URL.String(), header)
	if err != nil {
		return errors.Wrap(err, "can't dial to server")
	}
	defer resp.Body.Close()

	conn.SetCloseHandler(conn.CloseHandler())
	conn.SetPingHandler(conn.PingHandler())
	conn.SetPongHandler(conn.PongHandler())

	c.Conn = conn

	go c.listenRC()
	go c.listenWC()

	if err := c.auth(); err != nil {
		return errors.Wrap(err, "can't authorize")
	}

	return nil
}

func (c *Client) auth() error {
	if c.key != "" && c.secret != "" {
		exp := time.Now().Add(time.Minute).Unix()
		s, err := sign(c.secret, http.MethodGet+"/"+c.URL.RequestURI()+strconv.FormatInt(exp, 10))
		if err != nil {
			return errors.Wrap(err, "can't generate signature")
		}

		auth := WSEvent{
			Op: "authKeyExpires",
			Args: []interface{}{
				c.key,
				exp,
				s,
			},
		}

		if err := c.SendJSON(auth); err != nil {
			return errors.Wrap(err, "can't send json to server")
		}

	}
	return nil
}

func (c *Client) Listen() <-chan interface{} {
	var out = make(chan interface{})

	go func() {
		defer close(out)
		var val interface{}

		for v := range c.rc {
			var resp WSResponse
			if err := json.Unmarshal(v, &resp); err != nil {
				log.Println(err)
				break
			}

			bs, err := json.Marshal(resp.Data)
			if err != nil {
				log.Println(err)
				break
			}

			switch resp.Table {
			case Trade:
				var t []models.Trade
				if err := json.Unmarshal(bs, &t); err != nil {
					log.Printf("can't unmarshal %v", err)
					return
				}

				val = t
			case TradeBin1m, TradeBin5m, TradeBin1h, TradeBin1d:
				var t []models.TradeBucketed
				if err := json.Unmarshal(bs, &t); err != nil {
					log.Printf("can't unmarshal %v", err)
					return
				}
			case OrderBookL2, OrderBookL2_25, OrderBook10:
				var t []models.OrderBookL2
				if err := json.Unmarshal(bs, &t); err != nil {
					log.Printf("can't unmarshal %v", err)
					return
				}

				val = t
			case Order:
				var t []models.Order
				if err := json.Unmarshal(bs, &t); err != nil {
					log.Printf("can't unmarshal %v", err)
					return
				}

				val = t
			case Quote, QuoteBin1m, QuoteBin5m, QuoteBin1h, QuoteBin1d:
				var t []models.Quote
				if err := json.Unmarshal(bs, &t); err != nil {
					log.Printf("can't unmarshal %v", err)
					return
				}

				val = t
			default:
				val = bytes.TrimSpace(v)
			}

			out <- val
		}
	}()

	return out
}

func (c *Client) listenRC() error {
	defer close(c.rc)

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if err, ok := err.(*websocket.CloseError); ok {
				return errors.Wrapf(err, "error code %v", err.Code)
			}
			if _, ok := err.(*net.OpError); ok {
				return nil
			}

			return errors.Wrap(err, "")
		}

		c.rc <- message
	}

}

func (c *Client) listenWC() error {
	for v := range c.wc {
		w, err := c.Conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return errors.Wrap(err, "can't read response message")
		}
		if _, err := w.Write(v); err != nil {
			return errors.Wrap(err, "can't write to server")
		}
		if err := w.Close(); err != nil {
			return errors.Wrap(err, "can't close writer")
		}
	}
	return nil
}

func (c *Client) Disconnect() error {
	close(c.wc)
	return c.Conn.Close()
}

func sign(secret, message string) (string, error) {
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := h.Write([]byte(message)); err != nil {
		return "", errors.Wrapf(err, "can't write encrypted data")
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func (c *Client) SendJSON(v interface{}) error {
	bs, err := json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "")
	}
	c.wc <- bs
	return nil
}
