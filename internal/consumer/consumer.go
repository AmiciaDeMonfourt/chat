package consumer

import (
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
		"auto.offset.reset": "earliest",
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

func (c *Consumer) StartConsume() {
	defer c.consumer.Close()
	slog.Debug("start consume")
	for {
		ev := c.consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			slog.Debug("Consume message", "topic", e.TopicPartition.Topic, "val", e.Value, "key", e.Key)
		case kafka.Error:
			slog.Debug(e.Error(), "ctx", "consumer.StartConsume()")
		}
	}
}
