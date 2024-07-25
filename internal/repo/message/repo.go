package repo

import "test-service/internal/model"

type MessageRepository interface {
	SaveMessage(model.Message) (uint64, error)
	UpdateMessageStatus(model.Message) error
	GetMessages() ([]model.Message, error)
}
