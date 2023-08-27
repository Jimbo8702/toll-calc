package main

import (
	"time"

	"github.com/Jimbo8702/toll-calc/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err": err,
		}).Info("aggregate distance")
	}(time.Now())
	err = l.next.AggregateDistance(distance)
	return
}

func (l *LogMiddleware) CalculateInvoice(obuID int) (inv *types.Invoice, err error) {
	defer func(start time.Time) {
		var (
			distance float64
			amount	 float64
		)
		if inv != nil {
			distance = inv.TotalDistance
			amount = inv.TotalAmount
		}
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err": err,
			"obuID": obuID,
			"amount": amount,
			"distance": distance,
		}).Info("aggregate distance")
	}(time.Now())
	inv, err =l.next.CalculateInvoice(obuID)
	return
}
