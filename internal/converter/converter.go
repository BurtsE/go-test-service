package converter

import (
	"encoding/json"
	"test-service/internal/model"
)

func JsonToMessage(data []byte) (model.Message, error) {
	msg := model.Message{Status: model.InProgress}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func StatisticsToJson(unProcessed, processed uint) ([]byte, error) {
	stats := model.Statistics{
		Processed:   processed,
		UnProcessed: unProcessed,
	}
	data, err := json.Marshal(stats)
	if err != nil {
		return nil, err
	}
	return data, nil
}
