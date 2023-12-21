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

	req, err := http.NewRequest(http.MethodPost, c.Endpoint+"/aggregate", bytes.NewReader(b))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("service responded with unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *HTTPClient) GetInvoice(ctx context.Context, id int) (*types.Invoice, error) {
	invRequest := types.GetInvoiceRequest{
		ObuID: int32(id),
	}
	b, err := json.Marshal(&invRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.Endpoint+"/invoice", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("service responded with unexpected status code: %d", resp.StatusCode)
	}

	var inv types.Invoice
	if err := json.NewDecoder(resp.Body).Decode(&inv); err != nil {
		return nil, err
	}

	return &inv, nil
}
