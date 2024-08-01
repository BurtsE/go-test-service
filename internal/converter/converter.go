package converter

import (
	"encoding/json"
	"strconv"
	"test-service/internal/model"

	"github.com/segmentio/kafka-go"
)

func JsonToMessage(data []byte) (model.Message, error) {
	msg := model.Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func MessageToJson(msg model.Message) ([]byte, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func MessageToKafka(msg model.Message) kafka.Message {
	kafkaMsg := kafka.Message{
		Key:   []byte(strconv.FormatUint(msg.UUID, 10)),
		Value: []byte(msg.Text),
	}

	return kafkaMsg
}

func StatisticsToJson(stats model.Statistics) ([]byte, error) {
	data, err := json.Marshal(stats)
	if err != nil {
		return nil, err
	}
	return data, nil
}
