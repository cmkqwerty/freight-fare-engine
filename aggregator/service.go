package main

import (
	"fmt"
	"github.com/cmkqwerty/freight-fare-engine/types"
)

type Aggregator interface {
	AggregateDistance(distance types.Distance) error
}

type Storer interface {
	Insert(distance types.Distance) error
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	fmt.Println("processing and inserting distance:", distance)

	return i.store.Insert(distance)
}
