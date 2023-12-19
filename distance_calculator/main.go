package main

import "log"

const kafkaTopic = "obu-data"

func main() {
	var (
		svc CalculatorServicer
		err error
	)
	svc = NewCalculatorService()
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
