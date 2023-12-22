package aggservice

import (
	"context"
	"github.com/cmkqwerty/freight-fare-engine/types"
)

const basePrice = 3.15

type Service interface {
	Aggregate(ctx context.Context, distance types.Distance) error
	Calculate(ctx context.Context, id int) (*types.Invoice, error)
}

type Storer interface {
	Insert(distance types.Distance) error
	Get(id int) (float64, error)
}

type BasicService struct {
	store Storer
}

func newBasicService(store Storer) Service {
	return &BasicService{
		store: store,
	}
}

func (svc *BasicService) Aggregate(ctx context.Context, distance types.Distance) error {
	return svc.store.Insert(distance)
}

func (svc *BasicService) Calculate(ctx context.Context, id int) (*types.Invoice, error) {
	distance, err := svc.store.Get(id)
	if err != nil {
		return nil, err
	}

	invoice := &types.Invoice{
		OBUID:         id,
		TotalDistance: distance,
		TotalAmount:   basePrice * distance,
	}
	return invoice, nil
}

// NewAggregatorService returns a naive, stateless implementation of Service.
func NewAggregatorService() Service {
	var svc Service
	svc = newBasicService(NewMemoryStore())
	svc = newLoggingMiddleware()(svc)
	svc = newInstrumentationMiddleware()(svc)

	return svc
}
