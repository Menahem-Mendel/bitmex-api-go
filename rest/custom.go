package rest

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

type CustomService struct {
	request
}

func (c CustomService) Get(endpoint string, filters map[string]string) (interface{}, error) {
	var out interface{}
	u := make(url.Values)

	for k, v := range filters {
		u.Set(k, v)
	}

	bs, err := c.get(endpoint + "?" + u.Encode())
	if err != nil {
		return nil, errors.Wrapf(err, "can't get %s", endpoint)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal response body")
	}

	return out, nil
}

func (c CustomService) Post(endpoint string, data interface{}) (interface{}, error) {
	var out interface{}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal data")
	}

	bs, err := c.post(endpoint, body)
	if err != nil {
		return nil, errors.Wrapf(err, "can't post %s", endpoint)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal response body")
	}

	return out, nil
}

func (c CustomService) Put(endpoint string, data interface{}) (interface{}, error) {
	var out interface{}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal data")
	}

	bs, err := c.put(endpoint, body)
	if err != nil {
		return nil, errors.Wrapf(err, "can't put %s", endpoint)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal response body")
	}

	return out, nil
}

func (c CustomService) Delete(endpoint string, filters map[string]string) (interface{}, error) {
	var out interface{}
	u := make(url.Values)

	for k, v := range filters {
		u.Set(k, v)
	}

	bs, err := c.delete(endpoint + "?" + u.Encode())
	if err != nil {
		return nil, errors.Wrapf(err, "can't delete %s", endpoint)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal response body")
	}

	return out, nil
}
