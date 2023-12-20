package main

import "github.com/cmkqwerty/freight-fare-engine/types"

type GRPCAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewAggregatorGRPCServer(svc Aggregator) *GRPCAggregatorServer {
	return &GRPCAggregatorServer{
		svc: svc,
	}
}

func (g *GRPCAggregatorServer) AggregateDistance(request *types.AggregateRequest) error {
	distance := types.Distance{
		OBUID: int(request.ObuID),
		Value: request.Value,
		Unix:  request.Unix,
	}

	return g.svc.AggregateDistance(distance)
}
