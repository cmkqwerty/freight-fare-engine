package client

import (
	"context"
	"github.com/cmkqwerty/freight-fare-engine/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Endpoint string
	client   types.AggregatorClient
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := types.NewAggregatorClient(conn)

	return &GRPCClient{
		Endpoint: endpoint,
		client:   c,
	}, nil
}

func (c *GRPCClient) Aggregate(ctx context.Context, request *types.AggregateRequest) error {
	_, err := c.client.Aggregate(ctx, request)
	return err
}
