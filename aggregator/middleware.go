package main

import (
	"github.com/cmkqwerty/freight-fare-engine/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
	"time"
)

type MetricsMiddleware struct {
	errCounterAggregate prometheus.Counter
	errCounterCalculate prometheus.Counter
	reqCounterAggregate prometheus.Counter
	reqCounterCalculate prometheus.Counter
	reqLatencyAggregate prometheus.Histogram
	reqLatencyCalculate prometheus.Histogram
	next                Aggregator
}

func NewMetricsMiddleware(next Aggregator) *MetricsMiddleware {
	errCounterAggregate := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "aggregator_error_count",
		Name:      "aggregate",
	})
	errCounterCalculate := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "aggregator_error_count",
		Name:      "calculate",
	})
	reqCounterAggregate := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "aggregator_request_count",
		Name:      "aggregate",
	})
	reqCounterCalculate := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "aggregator_request_count",
		Name:      "calculate",
	})
	reqLatencyAggregate := promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "aggregator_request_latency",
		Name:      "aggregate",
		Buckets:   []float64{0.1, 0.5, 1, 2, 5, 10},
	})
	reqLatencyCalculate := promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "aggregator_request_latency",
		Name:      "calculate",
		Buckets:   []float64{0.1, 0.5, 1, 2, 5, 10},
	})

	return &MetricsMiddleware{
		errCounterAggregate: errCounterAggregate,
		errCounterCalculate: errCounterCalculate,
		reqCounterAggregate: reqCounterAggregate,
		reqCounterCalculate: reqCounterCalculate,
		reqLatencyAggregate: reqLatencyAggregate,
		reqLatencyCalculate: reqLatencyCalculate,
		next:                next,
	}
}

func (m *MetricsMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		m.reqCounterAggregate.Inc()
		m.reqLatencyAggregate.Observe(time.Since(start).Seconds())
		if err != nil {
			m.errCounterAggregate.Inc()
		}
	}(time.Now())
	err = m.next.AggregateDistance(distance)
	return
}

func (m *MetricsMiddleware) CalculateInvoice(id int) (invoice *types.Invoice, err error) {
	defer func(start time.Time) {
		m.reqCounterCalculate.Inc()
		m.reqLatencyCalculate.Observe(time.Since(start).Seconds())
		if err != nil {
			m.errCounterCalculate.Inc()
		}
	}(time.Now())
	invoice, err = m.next.CalculateInvoice(id)
	return
}

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":  time.Since(start),
			"error": err,
		}).Info("Aggregated Distance")
	}(time.Now())

	err = m.next.AggregateDistance(distance)
	return
}

func (m *LogMiddleware) CalculateInvoice(id int) (invoice *types.Invoice, err error) {
	defer func(start time.Time) {
		var (
			distance float64
			amount   float64
		)
		if invoice != nil {
			distance = invoice.TotalDistance
			amount = invoice.TotalAmount
		}

		logrus.WithFields(logrus.Fields{
			"took":     time.Since(start),
			"error":    err,
			"obuID":    id,
			"distance": distance,
			"amount":   amount,
		}).Info("Calculated Invoice")
	}(time.Now())

	invoice, err = m.next.CalculateInvoice(id)
	return
}
