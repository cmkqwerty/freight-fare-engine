package aggservice

import (
	"fmt"
	"github.com/cmkqwerty/freight-fare-engine/types"
)

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

func (s *MemoryStore) Get(id int) (float64, error) {
	distance, ok := s.data[id]
	if !ok {
		return 0, fmt.Errorf("no distance found for id %d", id)
	}

	return distance, nil
}
