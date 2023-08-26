package main

import (
	"log"

	"github.com/Jimbo8702/toll-calc/aggregator/client"
)

const (
	kafkaTopic = "obudata"
	aggregatorEndpoint = "http://127.0.0.1:4000/aggregate"
)

func main() {
	var (
		err error
		svc CalculatorServicer
		client = client.NewClient(aggregatorEndpoint)
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)
	
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}