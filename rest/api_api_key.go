package rest

import (
	"context"
	"fmt"
	"net/http"

	bitmex "github.com/Menahem-Mendel/bitmex-api-go"
	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/google/go-querystring/query"
)

// APIKeySnapshot
type APIKeySnapshot []models.APIKey

// APIKeyGetConf
type APIKeyGetConf struct {
	Reverse bool `url:"reverse,omitempty"`
}

// GetAPIKey
func (c Client) GetAPIKey(ctx context.Context, f APIKeyGetConf) (APIKeySnapshot, error) {
	var out APIKeySnapshot

	params, err := query.Values(f)
	if err != nil {
		return nil, fmt.Errorf("#Client.GetOrders query: %v", err)
	}

	x, err := c.Request(ctx, http.MethodGet, bitmex.APIKey+"?"+params.Encode(), nil, out)
	if err != nil {
		return nil, fmt.Errorf("#Client.GetOrders: %v", err)
	}

	switch x.(type) {
	case *APIKeySnapshot:
		out = *x.(*APIKeySnapshot)
	default:
		return nil, fmt.Errorf("#Client.GetOrders type %T isn't supported", x)
	}

	return out, nil
}
