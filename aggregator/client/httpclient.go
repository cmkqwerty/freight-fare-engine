package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cmkqwerty/freight-fare-engine/types"
	"net/http"
)

type HTTPClient struct {
	Endpoint string
}

func NewHTTPClient(endpoint string) *HTTPClient {
	return &HTTPClient{
		Endpoint: endpoint,
	}
}

func (c *HTTPClient) Aggregate(ctx context.Context, request *types.AggregateRequest) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.Endpoint, bytes.NewReader(b))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("service responded with unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
