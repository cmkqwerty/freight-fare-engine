package main

import (
	"context"
	"github.com/cmkqwerty/freight-fare-engine/types"
)

type GRPCAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewAggregatorGRPCServer(svc Aggregator) *GRPCAggregatorServer {
	return &GRPCAggregatorServer{
		svc: svc,
	}
}

func (g *GRPCAggregatorServer) Aggregate(ctx context.Context, request *types.AggregateRequest) (*types.None, error) {
	distance := types.Distance{
		OBUID: int(request.ObuID),
		Value: request.Value,
		Unix:  request.Unix,
	}

	return &types.None{}, g.svc.AggregateDistance(distance)
}
