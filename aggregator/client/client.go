package client

import (
	"context"
	"github.com/cmkqwerty/freight-fare-engine/types"
)

type Client interface {
	Aggregate(ctx context.Context, request *types.AggregateRequest) error
}
