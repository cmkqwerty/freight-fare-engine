package main

import (
	"encoding/json"
	"fmt"
	"github.com/cmkqwerty/freight-fare-engine/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type HTTPFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandlerFunc(h HTTPFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiErr, ok := err.(APIError); ok {
				writeJSON(w, apiErr.Code, map[string]string{"error": apiErr.Error()})
			}
		}
	}
}

type APIError struct {
	Code int
	Err  error
}

// Error implements error interface
func (e APIError) Error() string {
	return e.Err.Error()
}

func makeAPIError(code int, err error) APIError {
	return APIError{
		Code: code,
		Err:  err,
	}
}

type HTTPMetricHandler struct {
	reqCounter prometheus.Counter
	errCounter prometheus.Counter
	reqLatency prometheus.Histogram
}

func newHTTPMetricHandler(reqName string) *HTTPMetricHandler {
	return &HTTPMetricHandler{
		reqCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: fmt.Sprintf("http_%s_%s", reqName, "requests_total"),
			Name:      "aggregator",
		}),
		errCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: fmt.Sprintf("http_%s_%s", reqName, "errors_total"),
			Name:      "aggregator",
		}),
		reqLatency: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: fmt.Sprintf("http_%s_%s", reqName, "request_latency"),
			Name:      "aggregator",
			Buckets:   []float64{0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		}),
	}
}

func (mh *HTTPMetricHandler) instrument(h HTTPFunc) HTTPFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var err error
		defer func(start time.Time) {
			latency := time.Since(start).Seconds()
			logrus.WithFields(logrus.Fields{
				"latency": latency,
				"request": r.URL.Path,
				"err":     err,
			}).Info()
			mh.reqLatency.Observe(latency)
			mh.reqCounter.Inc()
			if err != nil {
				mh.errCounter.Inc()
			}
		}(time.Now())
		err = h(w, r)
		return err
	}
}

func handleGetInvoice(svc Aggregator) HTTPFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r.Method != http.MethodGet {
			return makeAPIError(http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		}

		obuID, err := strconv.Atoi(r.URL.Query().Get("obu"))
		if err != nil {
			return makeAPIError(http.StatusBadRequest, fmt.Errorf("invalid obu id"))
		}

		invoice, err := svc.CalculateInvoice(obuID)
		if err != nil {
			return makeAPIError(http.StatusInternalServerError, fmt.Errorf("failed to calculate invoice"))
		}

		return writeJSON(w, http.StatusOK, invoice)
	}
}

func handleAggregate(svc Aggregator) HTTPFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r.Method != http.MethodPost {
			return makeAPIError(http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		}

		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			return makeAPIError(http.StatusBadRequest, fmt.Errorf("invalid request body"))
		}

		if err := svc.AggregateDistance(distance); err != nil {
			return makeAPIError(http.StatusInternalServerError, fmt.Errorf("failed to aggregate distance"))
		}

		return writeJSON(w, http.StatusOK, nil)
	}
}
