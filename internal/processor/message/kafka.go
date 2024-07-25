package message

import (
	"context"
	"log"
	"test-service/internal/converter"
	"test-service/internal/model"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	reader *kafka.Reader
	writer *kafka.Writer
}

func NewKafka() (*Kafka, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		Topic:     "topic-A",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
		Topic:    "topic-B",
		Balancer: &kafka.LeastBytes{},
	}
	return &Kafka{
		reader: reader,
		writer: w,
	}, nil
}

func (k *Kafka) SendMessage(msg model.Message) error {
	log.Println("sending to kafka")
	message := converter.MessageToKafka(msg)
	err := k.writer.WriteMessages(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}
