package rest

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/pkg/errors"
)

type APIKeyService struct {
	request
}

// APIKeySnapshot
type APIKeySnapshot []*models.APIKey

// Get
func (a APIKeyService) Get(ctx context.Context, reverse bool) (APIKeySnapshot, error) {
	var out APIKeySnapshot

	params := make(url.Values)
	params.Set("timeout", strconv.FormatBool(reverse))

	uri := apiKey + "?" + params.Encode()

	bs, err := a.get(uri)
	if err != nil {
		return nil, errors.Wrapf(err, "can't get api keys (uri = %q)", uri)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}
