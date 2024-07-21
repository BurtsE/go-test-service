package service

import "test-service/internal/model"

type MessageService interface {
	SaveMessage() error
	GetMessages() ([]model.Message, error)
}
