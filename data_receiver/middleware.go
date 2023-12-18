package main

import (
	"github.com/cmkqwerty/freight-fare-engine/types"
	"github.com/sirupsen/logrus"
	"time"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (lm *LogMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID": data.OBUID,
			"lat":   data.Latitude,
			"lon":   data.Longitude,
			"took":  time.Since(start),
		}).Info("producing to kafka")
	}(time.Now())

	return lm.next.ProduceData(data)
}
