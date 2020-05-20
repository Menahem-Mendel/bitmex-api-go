package rest

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Menahem-Mendel/bitmex-api-go/models"

	bitmex "github.com/Menahem-Mendel/bitmex-api-go"
)

func (c Client) Do(ctx context.Context, method, uri string, data []byte) ([]byte, error) {
	req, err := c.NewRequest(ctx, method, uri, data)
	if err != nil {
		return nil, fmt.Errorf("#Client.Do req: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("#Client.Do resp: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		var errMsg models.ModelError
		if err := json.NewDecoder(resp.Body).Decode(&errMsg); err != nil {
			return nil, fmt.Errorf("#Client.Do decoding response error message: %v", err)
		}

		return nil, fmt.Errorf("#Client.Do %v: %v", errMsg.Error.Name, errMsg.Error.Message)
	}

	// -- rate limit sleep --
	if resp.Header.Get(bitmex.RemainingRequestNum) == "1" {
		num, err := strconv.ParseInt(resp.Header.Get(bitmex.LimitReset), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("#Client.Do parseint: %v", err)
		}

		log.Println("wait untill rate limit resets...")
		time.Sleep(time.Until(time.Unix(num, 0)))
	}

	return ioutil.ReadAll(resp.Body)
}

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

	fmt.Println(req.Method, req.URL.Path, req.Proto)
	fmt.Println("Host:", req.URL.Host)
	fmt.Println("Header[Content-Type] =", req.Header.Get("content-type"))
	fmt.Println("Header[Accept] =", req.Header.Get("accept"))
	fmt.Println("Header["+bitmex.KeyAPI+"] =", req.Header.Get(bitmex.KeyAPI))
	fmt.Println("Header["+bitmex.ExpiresAPI+"] =", req.Header.Get(bitmex.ExpiresAPI))
	fmt.Println("Header["+bitmex.SignatureAPI+"] =", req.Header.Get(bitmex.SignatureAPI))
	fmt.Println("User-Agent:", req.UserAgent())
	fmt.Println()

	return req, nil
}

func signature(secret, method, uri, expires, data string) (string, error) {
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := h.Write([]byte(method + uri + expires + data)); err != nil {
		return "", fmt.Errorf("#signature %v", err)
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
