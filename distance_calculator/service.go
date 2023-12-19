package main

import (
	"github.com/cmkqwerty/freight-fare-engine/types"
	"math"
)

type CalculatorServicer interface {
	CalculateDistance(data types.OBUData) (float64, error)
}

type CalculatorService struct {
	points [][]float64
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{
		points: make([][]float64, 0),
	}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0
	if len(s.points) > 0 {
		lastPoint := s.points[len(s.points)-1]
		distance := calculateDistance(lastPoint[0], lastPoint[1], data.Latitude, data.Longitude)
		return distance, nil
	}

	s.points = append(s.points, []float64{data.Latitude, data.Longitude})
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
