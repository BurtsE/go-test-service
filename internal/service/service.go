package service

import "test-service/internal/model"

type MessageService interface {
	SaveMessage(model.Message) error
	GetStatistics() (uint, uint, error)
}
