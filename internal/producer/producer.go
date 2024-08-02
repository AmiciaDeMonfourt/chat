package producer

import (
	"bytes"
	"log/slog"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	kafkaProducer *kafka.Producer
}

func New() *Producer {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		slog.Error("failed to create producer: "+err.Error(), "ctx", "producer.New()")
		os.Exit(1)
	}

	return &Producer{
		kafkaProducer: producer,
	}
}

func (p *Producer) WriteAsync() error {
	topic := "test-topic"

	msg := &kafka.Message{
		Value: bytes.NewBufferString("value").Bytes(),
		Key:   bytes.NewBufferString("key").Bytes(),

		TopicPartition: kafka.TopicPartition{
			Topic: &topic,
		},
	}

	return p.kafkaProducer.Produce(msg, nil)
}
