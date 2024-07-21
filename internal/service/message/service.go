package message

import (
	repo "test-service/internal/repo/message"
	def "test-service/internal/service"
)

var _ def.MessageService = (*service)(nil)

type service struct {
	messageRepo repo.MessageRepository
}

func NewService(messageRepository *repo.MessageRepository) *service {
	return &service{
		messageRepo: *messageRepository,
	}
}
