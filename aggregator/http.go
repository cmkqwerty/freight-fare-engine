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

type HTTPMetricHandler struct {
	reqCounter prometheus.Counter
	reqLatency prometheus.Histogram
}

func newHTTPMetricHandler(reqName string) *HTTPMetricHandler {
	return &HTTPMetricHandler{
		reqCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: fmt.Sprintf("http_%s_%s", reqName, "requests_total"),
			Name:      "aggregator",
		}),
		reqLatency: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: fmt.Sprintf("http_%s_%s", reqName, "request_latency"),
			Name:      "aggregator",
			Buckets:   []float64{0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		}),
	}
}

func (mh *HTTPMetricHandler) instrument(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time) {
			latency := time.Since(start).Seconds()
			logrus.WithFields(logrus.Fields{
				"latency": latency,
				"request": r.URL.Path,
			}).Info()
			mh.reqLatency.Observe(latency)
		}(time.Now())
		mh.reqCounter.Inc()
		h(w, r)
	}
}

func handleGetInvoice(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}

		obuID, err := strconv.Atoi(r.URL.Query().Get("obu"))
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid obu id"})
			return
		}

		invoice, err := svc.CalculateInvoice(obuID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, invoice)
	}
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}

		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}
}
