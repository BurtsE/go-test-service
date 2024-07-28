package message

import (
	"context"
	"test-service/internal/config"
	"test-service/internal/converter"
	"test-service/internal/model"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	writer *kafka.Writer
}

func NewKafka(cfg *config.Config) (*Kafka, error) {

	w := &kafka.Writer{
		Addr:     kafka.TCP(cfg.Brokers...),
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &Kafka{
		writer: w,
	}, nil
}

func (k *Kafka) SendMessage(msg model.Message) error {
	message := converter.MessageToKafka(msg)
	err := k.writer.WriteMessages(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}
