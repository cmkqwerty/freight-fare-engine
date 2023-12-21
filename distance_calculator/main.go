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

	httpClient := client.NewHTTPClient(aggregatorEndpoint)
	/*
		grpcClient, err := client.NewGRPCClient(aggregatorEndpoint)
		if err != nil {
			log.Fatal(err)
		}
	*/

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, httpClient)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
