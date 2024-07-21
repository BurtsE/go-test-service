package repo

import "test-service/internal/model"

type MessageRepository interface {
	SaveMessage(model.Message) error
	GetMessages() ([]model.Message, error)
}
