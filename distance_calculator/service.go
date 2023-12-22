package main

import (
	"github.com/cmkqwerty/freight-fare-engine/types"
	"math"
)

type CalculatorServicer interface {
	CalculateDistance(data types.OBUData) (float64, error)
}

type CalculatorService struct {
	lastPoint []float64
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0
	if len(s.lastPoint) > 0 {
		distance := calculateDistance(s.lastPoint[0], s.lastPoint[1], data.Latitude, data.Longitude)
		return distance, nil
	}

	s.lastPoint = []float64{data.Latitude, data.Longitude}
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
