package message

import (
	"test-service/internal/config"
	repo "test-service/internal/repo/message"
	def "test-service/internal/service"
)

var _ def.MessageService = (*service)(nil)

type service struct {
	messageRepo repo.MessageRepository
}

func NewService(messageRepository repo.MessageRepository, cfg *config.Config) *service {
	return &service{
		messageRepo: messageRepository,
	}
}
