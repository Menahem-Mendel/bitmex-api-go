package rest

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Transport struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	httpDo     func(c *http.Client, req *http.Request) (*http.Response, error)
}

func (t Transport) Exec(req *Request) ([]byte, error) {
	var data []byte

	body := bytes.NewReader(req.Data)

	u, err := t.BaseURL.Parse(req.URI)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	r, err := http.NewRequest(req.Method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	for k, v := range req.Headers {
		r.Header.Add(k, v)
	}

	data, err = t.do(r)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return data, nil
}

func (t Transport) do(req *http.Request) ([]byte, error) {
	resp, err := t.httpDo(t.HTTPClient, req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()

	r := NewResponse(resp)

	if r.ErrorBody != nil {
		log.Println(string(r.ErrorBody))
	}

	return r.Body, nil
}
