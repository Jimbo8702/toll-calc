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