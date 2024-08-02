package consumer

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	consumer *kafka.Consumer
}

func New(topics []string) *Consumer {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "lohi",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		slog.Error("failed to create consumer: "+err.Error(), "ctx", "consumer.New()")
		os.Exit(1)
	}

	if err := consumer.SubscribeTopics(topics, nil); err != nil {
		slog.Error("failed to subscribe on topics: "+err.Error(), "ctx", "consumer.New()")
		os.Exit(1)
	}

	return &Consumer{
		consumer: consumer,
	}
}

func (c *Consumer) Consume() {
	for {
		kafkaEvent := c.consumer.Poll(100)

		switch event := kafkaEvent.(type) {
		case *kafka.Message:
			fmt.Println(event.Value)
		}
	}
}
