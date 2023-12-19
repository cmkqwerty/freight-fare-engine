package main

import "github.com/cmkqwerty/freight-fare-engine/types"

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (s *MemoryStore) Insert(distance types.Distance) error {
	s.data[distance.OBUID] += distance.Value

	return nil
}
