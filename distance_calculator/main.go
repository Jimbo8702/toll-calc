package main

import (
	"log"

	"github.com/Jimbo8702/toll-calc/aggregator/client"
)

const (
	kafkaTopic = "obudata"
	aggregatorEndpoint = "http://127.0.0.1:4000"
)

func main() {
	var (
		err error
		svc CalculatorServicer
		client = client.NewHTTPClient(aggregatorEndpoint)
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)
	
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}