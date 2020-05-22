package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/google/go-querystring/query"
)

type APIKeyService struct {
	RequestFactory
	Synchronous
}

// APIKeySnapshot
type APIKeySnapshot []models.APIKey

// APIKeyGetConf
type APIKeyGetConf struct {
	Reverse bool `url:"reverse,omitempty"`
}

// Get
func (a *APIKeyService) Get(ctx context.Context, f APIKeyGetConf) (*APIKeySnapshot, error) {
	var out *APIKeySnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req := NewRequest(http.MethodGet, Order+params.Encode(), nil)

	bs, err := a.Exec(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return out, nil
}
