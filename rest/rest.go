package rest

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	bitmex "github.com/Menahem-Mendel/bitmex-api-go"
)

func get(ctx context.Context, path *url.URL, key string, expires int64) ([]byte, error) {
	var out []byte

	req, err := request(ctx, path, http.MethodGet, key, strconv.FormatInt(expires, 10), "")
	if err != nil {
		return out, fmt.Errorf("#Client.get req: %v", err)
	}

	resp, err := do(req)
	if err != nil {
		return out, fmt.Errorf("#Client.get resp: %v", err)
	}
	defer resp.Body.Close()

	//!try to change it to effisient way (like json decoder)
	out, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("#Client.get reading response: %v", err)
	}

	return out, nil
}

// func get(ctx context.Context, base url.URL, key, endpoint string, f interface{}, in interface{}) (interface{}, error) {
// 	out := reflect.ValueOf(in).Interface()
// 	// out := reflect.New(reflect.TypeOf(in)).Interface()

// 	params, err := query.Values(f)
// 	if err != nil {
// 		return nil, fmt.Errorf("#get query: %v", err)
// 	}

// 	path, err := base.Parse(tradeBucketed + "?" + params.Encode())
// 	if err != nil {
// 		return nil, fmt.Errorf("#get path: %v", err)
// 	}

// 	req, err := request(ctx, http.MethodGet, key, path, "")
// 	if err != nil {
// 		return nil, fmt.Errorf("#get req: %v", err)
// 	}

// 	resp, err := do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("#get resp: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
// 		return nil, fmt.Errorf("#get unmarshal: %v", err)
// 	}

// 	return out, nil

// }

func request(ctx context.Context, url *url.URL, method, key, expires, data string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("#request: %v", err)
	}

	if key != "" {
		if err := setHeaders(req, key, method, url.RequestURI(), expires, data); err != nil {
			return nil, fmt.Errorf("#request setHeaders: %v", err)
		}
	}

	return req, nil
}

func setHeaders(req *http.Request, key, method, path, expires, data string) error {

	req.Header.Set("content-type", "application/json; charset=utf-8")

	if v, ok := req.Context().Value(bitmex.ContextAPIKey).(string); ok {
		req.Header.Set(bitmex.SignatureAPI, signature(v, method, path, expires, data))
		req.Header.Set(bitmex.KeyAPI, key)
		req.Header.Set(bitmex.ExpiresAPI, expires)
	}

	if err := req.Context().Err(); err != nil {
		return fmt.Errorf("#setHeaders could not authoarize: %v", err)
	}

	return nil
}

func do(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("#do resp: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		errMessage, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("#do err reading error message from response")
		}

		return nil, fmt.Errorf("#do got status: %v\n%s", resp.Status, errMessage)
	}

	// check rate limit, sleep if reached limit
	if resp.Header.Get(bitmex.RemainingRequestNum) == "1" {
		num, err := strconv.ParseInt(resp.Header.Get(bitmex.LimitReset), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("#do parseint: %v", err)
		}

		log.Println("wait untill rate limit resets...")
		time.Sleep(time.Until(time.Unix(num, 0)))
	}

	return resp, nil
}

func signature(secret, method, requestURI, expires, data string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(method + requestURI + expires + data))

	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
