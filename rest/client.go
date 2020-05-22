package rest

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	bitmex "github.com/Menahem-Mendel/bitmex-api-go"
)

// Client stores the requiared data for requesting
type Client struct {
	apikey string

	Orders  OrderService
	APIKeys APIKeyService
	Trades  TradeService

	Synchronous
}

// NewClient creates new client for requesting data from the bitmex exchange
func NewTestClient() *Client {
	httpDo := func(c *http.Client, req *http.Request) (*http.Response, error) {
		return c.Do(req)
	}
	return newClient(bitmex.BitmexHostTestnet, httpDo)
}

// NewAuthClient return authoarized client
// it takes public key and expires time in unix format
// the secret key you need to pass with requesting as a context value
func NewClient() *Client {
	httpDo := func(c *http.Client, req *http.Request) (*http.Response, error) {
		return c.Do(req)
	}

	return newClient(bitmex.BitmexHost, httpDo)
}

func (c *Client) Credentials(apikey string) *Client {
	c.apikey = apikey
	return c
}

func newClient(base string, httpDo func(*http.Client, *http.Request) (*http.Response, error)) *Client {

	baseURL := &url.URL{
		Scheme: "https",
		Host:   base,
		Path:   bitmex.BitmexAPIV1,
	}

	sync := &Transport{
		BaseURL:    baseURL,
		httpDo:     httpDo,
		HTTPClient: http.DefaultClient,
	}

	c := &Client{
		Synchronous: sync,
	}

	c.Orders = OrderService{Synchronous: c, RequestFactory: c}
	c.APIKeys = APIKeyService{Synchronous: c, RequestFactory: c}
	c.Trades = TradeService{Synchronous: c, RequestFactory: c}

	return c
}

func (c *Client) NewRequest(ctx context.Context, method, uri string, data []byte) *Request {
	req := NewRequest(method, uri, data)
	req.ctx = ctx

	secret, ok := ctx.Value(bitmex.ContextAPIKey).(string)
	if !ok {
		log.Panicln("secret isn't valid")
	}

	expires := strconv.FormatInt(time.Now().Add(time.Minute).Unix(), 10)

	sign := signature(secret, method, uri, expires, string(data))

	req.Headers["api-signature"] = sign
	req.Headers["api-key"] = c.apikey
	req.Headers["api-expires"] = expires

	return req
}

func signature(secret, method, uri, expires, data string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(method + bitmex.BitmexAPIV1 + uri + expires + data))

	return hex.EncodeToString(h.Sum(nil))
}
