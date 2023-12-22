package main

import (
	"fmt"
	"github.com/cmkqwerty/freight-fare-engine/aggregator/client"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const kafkaTopic = "obu-data"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var (
		aggregatorEndpoint string
		svc                CalculatorServicer
		err                error
	)
	aggregatorEndpoint = fmt.Sprintf("http://%s", os.Getenv("AGGREGATE_HTTP_ENDPOINT"))
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
