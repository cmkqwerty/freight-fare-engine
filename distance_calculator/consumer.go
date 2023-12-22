package main

import (
	"context"
	"encoding/json"
	"github.com/cmkqwerty/freight-fare-engine/aggregator/client"
	"github.com/cmkqwerty/freight-fare-engine/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"time"
)

type KafkaConsumer struct {
	consumer         *kafka.Consumer
	isRunning        bool
	calcService      CalculatorServicer
	aggregatorClient client.Client
}

func NewKafkaConsumer(topic string, svc CalculatorServicer, aggregatorClient client.Client) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	c.SubscribeTopics([]string{topic}, nil)

	return &KafkaConsumer{
		consumer:         c,
		calcService:      svc,
		aggregatorClient: aggregatorClient,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("Starting Kafka transport")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) Stop() {
	logrus.Info("Stopping Kafka transport")
	c.isRunning = false
	c.consumer.Close()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("Kafka consumer error: %v (%v)\n", err, msg)
			continue
		}

		var data types.OBUData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("Error JSON marshalling OBUData: %v", err)
			logrus.WithFields(logrus.Fields{
				"err":       err,
				"requestID": data.RequestID,
			})
			continue
		}

		distance, err := c.calcService.CalculateDistance(data)
		if err != nil {
			logrus.Errorf("Error calculating distance: %v", err)
			continue
		}

		req := &types.AggregateRequest{
			ObuID: int32(data.OBUID),
			Value: distance,
			Unix:  time.Now().UnixNano(),
		}
		if err := c.aggregatorClient.Aggregate(context.Background(), req); err != nil {
			logrus.Errorf("Error aggregating invoice: %v", err)
			continue
		}
	}
}
