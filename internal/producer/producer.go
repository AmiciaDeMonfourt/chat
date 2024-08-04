package producer

import (
	"log/slog"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	kafkaProducer *kafka.Producer
	topic         *string
}

func New(topic string) *Producer {
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
		topic:         &topic,
	}
}

func (p *Producer) StartProduce() {
	defer p.kafkaProducer.Close()
	for e := range p.kafkaProducer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				slog.Error(ev.TopicPartition.Error.Error(), "ctx", "producer.StartProduce()")
				continue
			}

			slog.Debug("Sucessfully produced record",
				"topic", *ev.TopicPartition.Topic,
				"partition", ev.TopicPartition.Partition,
				"offset", ev.TopicPartition.Offset,
				"msg", string(ev.Value),
			)
		}
	}
}

func (p *Producer) Write(key, value []byte) error {
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: p.topic,
		},
		Key:   key,
		Value: value,
	}

	return p.kafkaProducer.Produce(msg, nil)
}
