package main

import (
	"github.com/cmkqwerty/freight-fare-engine/aggregator/client"
	"log"
)

const (
	kafkaTopic         = "obu-data"
	aggregatorEndpoint = "http://localhost:3000/aggregate"
)

func main() {
	var (
		svc CalculatorServicer
		err error
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client.NewClient(aggregatorEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
