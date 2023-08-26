package main

import (
	"github.com/Jimbo8702/toll-calc/types"
)

type Aggregator interface {
	AggregateDistance(types.Distance) error
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
	// fmt.Println("Processing and inserting distance in the storage:", distance)
	return i.store.Insert(distance) 
}