package aggendpoint

import (
	"context"
	"github.com/cmkqwerty/freight-fare-engine/gokitimp/aggsvc/aggservice"
	"github.com/cmkqwerty/freight-fare-engine/types"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/ratelimit"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"time"
)

type Set struct {
	AggregateEndpoint endpoint.Endpoint
	CalculateEndpoint endpoint.Endpoint
}

func New(svc aggservice.Service, logger log.Logger) Set {
	duration := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "freight_fare_engine",
		Subsystem: "aggservice",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds.",
	}, []string{"method", "success"})

	var aggregateEndpoint endpoint.Endpoint
	{
		aggregateEndpoint = MakeAggregateEndpoint(svc)

		aggregateEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(aggregateEndpoint)
		aggregateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(aggregateEndpoint)
		aggregateEndpoint = LoggingMiddleware(log.With(logger, "method", "Aggregate"))(aggregateEndpoint)
		aggregateEndpoint = InstrumentingMiddleware(duration.With("method", "Aggregate"))(aggregateEndpoint)
	}

	var calculateEndpoint endpoint.Endpoint
	{
		calculateEndpoint = MakeCalculateEndpoint(svc)

		calculateEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(calculateEndpoint)
		calculateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(calculateEndpoint)
		calculateEndpoint = LoggingMiddleware(log.With(logger, "method", "Invoice"))(calculateEndpoint)
		calculateEndpoint = InstrumentingMiddleware(duration.With("method", "Invoice"))(calculateEndpoint)
	}

	return Set{
		AggregateEndpoint: aggregateEndpoint,
		CalculateEndpoint: calculateEndpoint,
	}
}

type AggregateRequest struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obuID"`
	Unix  int64   `json:"unix"`
}

type AggregateResponse struct {
	Err error `json:"err,omitempty"`
}

func (s Set) Aggregate(ctx context.Context, distance types.Distance) error {
	_, err := s.AggregateEndpoint(ctx, AggregateRequest{
		Value: distance.Value,
		OBUID: distance.OBUID,
		Unix:  distance.Unix,
	})
	return err
}

type CalculateRequest struct {
	OBUID int `json:"obuID"`
}

type CalculateResponse struct {
	OBUID         int     `json:"obuID"`
	TotalDistance float64 `json:"totalDistance"`
	TotalAmount   float64 `json:"totalAmount"`
	Err           error   `json:"err,omitempty"`
}

func (s Set) Calculate(ctx context.Context, id int) (*types.Invoice, error) {
	resp, err := s.CalculateEndpoint(ctx, CalculateRequest{OBUID: id})
	if err != nil {
		return nil, err
	}

	result := resp.(CalculateResponse)
	return &types.Invoice{
		OBUID:         result.OBUID,
		TotalDistance: result.TotalDistance,
		TotalAmount:   result.TotalAmount,
	}, nil
}

func MakeAggregateEndpoint(svc aggservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AggregateRequest)
		err = svc.Aggregate(ctx, types.Distance{
			OBUID: req.OBUID,
			Value: req.Value,
			Unix:  req.Unix,
		})

		return AggregateResponse{Err: err}, nil
	}
}

func MakeCalculateEndpoint(svc aggservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CalculateRequest)
		invoice, err := svc.Calculate(ctx, req.OBUID)

		return CalculateResponse{
			Err:           err,
			OBUID:         invoice.OBUID,
			TotalDistance: invoice.TotalDistance,
			TotalAmount:   invoice.TotalAmount,
		}, nil
	}
}
