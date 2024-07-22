package converter

import (
	"encoding/json"
	"test-service/internal/model"
)

func JsonToMessage(data []byte) error {
	msg := model.Message{Status: model.InProgress}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return err
	}
	return nil
}
