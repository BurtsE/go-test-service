package repo

import "test-service/internal/model"

type MessageRepository interface {
	SaveMessage() error
	GetMessages() ([]model.Message, error)
}
