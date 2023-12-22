package aggservice

import (
	"context"
	"github.com/cmkqwerty/freight-fare-engine/types"
)

type Middleware func(Service) Service

type loggingMiddleware struct {
	next Service
}

func newLoggingMiddleware() Middleware {
	return func(next Service) Service {
		return loggingMiddleware{
			next: next,
		}
	}
}

func (mw loggingMiddleware) Aggregate(ctx context.Context, distance types.Distance) error {
	return mw.next.Aggregate(ctx, distance)
}

func (mw loggingMiddleware) Calculate(ctx context.Context, id int) (*types.Invoice, error) {
	return mw.next.Calculate(ctx, id)
}

type instrumentationMiddleware struct {
	next Service
}

func newInstrumentationMiddleware() Middleware {
	return func(next Service) Service {
		return instrumentationMiddleware{
			next: next,
		}
	}
}

func (mw instrumentationMiddleware) Aggregate(ctx context.Context, distance types.Distance) error {
	return mw.next.Aggregate(ctx, distance)
}

func (mw instrumentationMiddleware) Calculate(ctx context.Context, id int) (*types.Invoice, error) {
	return mw.next.Calculate(ctx, id)
}
