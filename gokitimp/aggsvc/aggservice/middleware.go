package aggservice

import (
	"context"
	"github.com/cmkqwerty/freight-fare-engine/types"
	"github.com/go-kit/log"
	"time"
)

type Middleware func(Service) Service

type loggingMiddleware struct {
	log  log.Logger
	next Service
}

func newLoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{
			next: next,
			log:  logger,
		}
	}
}

func (mw loggingMiddleware) Aggregate(ctx context.Context, distance types.Distance) (err error) {
	defer func(start time.Time) {
		mw.log.Log("method", "Aggregate", "took", time.Since(start), "OBU", distance.OBUID, "distance", distance.Value, "error", err)
	}(time.Now())
	err = mw.next.Aggregate(ctx, distance)
	return
}

func (mw loggingMiddleware) Calculate(ctx context.Context, id int) (invoice *types.Invoice, err error) {
	defer func(start time.Time) {
		mw.log.Log("method", "Calculate", "took", time.Since(start), "id", id, "error", err)
	}(time.Now())
	invoice, err = mw.next.Calculate(ctx, id)
	return
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
