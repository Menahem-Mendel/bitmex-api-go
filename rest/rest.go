package rest

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	bitmex "github.com/Menahem-Mendel/bitmex-api-go"
	"github.com/Menahem-Mendel/bitmex-api-go/models"
)

// Request
func (c Client) Request(ctx context.Context, method, uri string, data []byte, in interface{}) (interface{}, error) {
	var out = reflect.New(reflect.TypeOf(in)).Interface()

	req, err := c.NewRequest(ctx, method, uri, data)
	if err != nil {
		return nil, fmt.Errorf("#Client.Request req: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("#Client.Request resp: %v", err)
	}
	defer resp.Body.Close()

	// -- rate limit sleep --
	if remain, err := strconv.ParseInt(resp.Header.Get(bitmex.RemainingRequestNum), 10, 32); err != nil {
		return nil, fmt.Errorf("#Client.Request: %v", err)
	} else if remain < 1 {
		t, err := strconv.ParseInt(resp.Header.Get(bitmex.LimitReset), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("#Client.Request parseint: %v", err)
		}

		log.Println("wait untill rate limit resets...")
		time.Sleep(time.Until(time.Unix(t, 0)))
	}

	if resp.StatusCode != http.StatusOK {
		var errMsg models.ModelError
		if err := json.NewDecoder(resp.Body).Decode(&errMsg); err != nil {
			return nil, fmt.Errorf("#Client.Request decoding response error message: %v", err)
		}

		return nil, fmt.Errorf("#Client.Request status %v error %v: %v", resp.Status, errMsg.Error.Name, errMsg.Error.Message)
	}

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("#Client.Request decoding response error message: %v", err)
	}

	return out, nil
}

// NewRequest
func (c Client) NewRequest(ctx context.Context, method, uri string, data []byte) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.Base.String()+uri, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("#Client.NewRequest request: %v", err)
	}

	v, ok := req.Context().Value(bitmex.ContextAPIKey).(string)
	if !ok {
		return nil, fmt.Errorf("#Client.NewRequest context value type should be string")
	}
	if err := req.Context().Err(); err != nil {
		return nil, fmt.Errorf("#Client.NewRequest context: %v", err)
	}

	expires := strconv.FormatInt(time.Now().Add(time.Minute).Unix(), 10)
	sign, err := signature(v, method, c.Base.Path+uri, expires, string(data))
	if err != nil {
		return nil, fmt.Errorf("#Client.NewRequest signature: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Jpro")
	req.Header.Set(bitmex.SignatureAPI, sign)
	req.Header.Set(bitmex.KeyAPI, c.key)
	req.Header.Set(bitmex.ExpiresAPI, expires)

	fmt.Println("------------------  REQUEST  ------------------")
	fmt.Println(req.Method, req.URL.Path, req.Proto)
	fmt.Println("Host:", req.URL.Host)
	for k, v := range req.Header {
		fmt.Printf("Header[%v] = %v\n", k, v)
	}
	fmt.Println()

	return req, nil
}

// singnature
func signature(secret, method, uri, expires, data string) (string, error) {
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := h.Write([]byte(method + uri + expires + data)); err != nil {
		return "", fmt.Errorf("#signature %v", err)
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
