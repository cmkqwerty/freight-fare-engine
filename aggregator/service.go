package main

import (
	"fmt"
	"github.com/cmkqwerty/freight-fare-engine/types"
)

const basePrice = 3.15

type Aggregator interface {
	AggregateDistance(distance types.Distance) error
	CalculateInvoice(id int) (*types.Invoice, error)
}

type Storer interface {
	Insert(distance types.Distance) error
	Get(id int) (float64, error)
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

func (i *InvoiceAggregator) CalculateInvoice(id int) (*types.Invoice, error) {
	distance, err := i.store.Get(id)
	if err != nil {
		return nil, err
	}

	invoice := &types.Invoice{
		OBUID:         id,
		TotalDistance: distance,
		TotalAmount:   basePrice * distance,
	}

	return invoice, nil
}
