package main

import (
	"github.com/cmkqwerty/freight-fare-engine/types"
	"github.com/sirupsen/logrus"
	"time"
)

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
		}).Info("Aggregated distance")
	}(time.Now())

	err = m.next.AggregateDistance(distance)
	return
}
