package main

import "github.com/cmkqwerty/freight-fare-engine/types"

type MemoryStore struct{}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) Insert(distance types.Distance) error {
	return nil
}
