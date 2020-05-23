package rest

import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	bitmex "github.com/Menahem-Mendel/bitmex-api-go"
	"github.com/pkg/errors"
)

// headers
const (
	remaining = "x-ratelimit-remaining"
	limit     = "x-ratelimit-limit"
	reset     = "x-ratelimit-reset"
)

type request interface {
	get(uri string) ([]byte, error)
	post(uri string, body []byte) ([]byte, error)
	put(uri string, body []byte) ([]byte, error)
	delete(uri string) ([]byte, error)
}

type Client struct {
	URL url.URL

	key    string
	secret string

	Orders  OrderService
	APIKeys APIKeyService

	Trades TradeService
	Book   BookService
	Stats  StatsService

	Custom CustomService

	// Users UserService
	// Settlement SettlementService
	// Leaderboard LeaderboardService
	// Quotes QuoteService
	// Positions PositionService
}

func NewClientTestnet() *Client {
	return newClient(bitmex.BitmexHostTestnet)
}

func NewClient() *Client {
	return newClient(bitmex.BitmexHost)
}

func (c *Client) Auth(secret, key string) *Client {
	c.key = key
	c.secret = secret
	return c
}

func newClient(base string) *Client {
	c := &Client{
		URL: url.URL{
			Scheme: "https",
			Host:   base,
			Path:   bitmex.BitmexAPIV1,
		},
	}

	c.Orders = OrderService{request: c}
	c.APIKeys = APIKeyService{request: c}
	c.Trades = TradeService{request: c}
	c.Book = BookService{request: c}
	c.Stats = StatsService{request: c}
	c.Custom = CustomService{request: c}

	return c
}

func (c Client) get(uri string) ([]byte, error) {
	return c.request(http.MethodGet, uri, nil)
}

// Put sends PUT request to the server
func (c Client) put(uri string, body []byte) ([]byte, error) {
	return c.request(http.MethodPut, uri, body)
}

// Post sends POST post request to the server
func (c Client) post(uri string, body []byte) ([]byte, error) {
	return c.request(http.MethodPost, uri, body)
}

// Delete sends DELETE request to the server
func (c Client) delete(uri string) ([]byte, error) {
	return c.request(http.MethodDelete, uri, nil)
}

func (c Client) request(method, uri string, body []byte) ([]byte, error) {
	u, err := c.URL.Parse(uri)
	if err != nil {
		return nil, errors.Wrapf(err, "can't parse (uri = %q)", uri)
	}

	log.Printf("INFO - Requesting %q", u.String())
	req, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "can't make new request")
	}

	setHeaders(req)

	// authorization not required
	if c.secret != "" && c.key != "" {
		log.Printf("INFO - Authorizing via API key")
		exp := strconv.FormatInt(time.Now().Add(time.Minute).Unix(), 10)

		s, err := sign(c.secret, method, uri, exp, string(body))
		if err != nil {
			return nil, errors.Wrap(err, "can't generate signature")
		}

		req.Header.Set("api-signature", s)
		req.Header.Set("api-key", c.key)
		req.Header.Set("api-expires", exp)
	}

	return do(req)
}

func setHeaders(req *http.Request) {
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("user-agent", "Jpro")
}

func do(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't do request")
	}
	defer resp.Body.Close()

	if err := check(resp); err != nil {
		return nil, errors.Wrapf(err, "response status code %s", resp.Status)
	}

	return scan(resp.Body)
}

func check(resp *http.Response) error {
	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusIMUsed {
		rmn, err := strconv.Atoi(resp.Header.Get(remaining))
		if err != nil {
			return errors.Wrapf(err, "can't convert string (%s) to int", resp.Header.Get(remaining))
		} else if rmn < 1 {
			rst, err := strconv.ParseInt(resp.Header.Get(reset), 10, 64)
			if err != nil {
				return errors.Wrapf(err, "can't convert string (%s) to int", resp.Header.Get(remaining))
			}
			log.Printf("SLEEP - untill rate limit resets %v", time.Until(time.Unix(int64(rst), 0)))
			time.Sleep(time.Until(time.Unix(int64(rst), 0)))
		}

		bs, err := scan(resp.Body)
		if err != nil {
			return errors.Wrap(err, "can't scan response body")
		}
		return errors.Errorf("%s", bs)
	}

	return nil
}

func scan(resp io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 0, 64*1024)

	s := bufio.NewScanner(resp)
	s.Buffer(buf, 1024*1024)
	for s.Scan() {
		out = s.Bytes()
	}
	if err := s.Err(); err != nil {
		return nil, errors.Wrap(err, "can't scan response body")
	}

	return out, nil
}

func sign(secret, method, uri, expires, data string) (string, error) {
	msg := method + bitmex.BitmexAPIV1 + uri + expires + data

	h := hmac.New(sha256.New, []byte(secret))
	if _, err := h.Write([]byte(msg)); err != nil {
		return "", errors.Wrapf(err, "can't write encrypted data")
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
