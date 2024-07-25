package processor

import "test-service/internal/model"

type MessageProcessor interface {
	SendMessage(model.Message) error
}
