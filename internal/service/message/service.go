package message

import (
	"test-service/internal/config"
	"test-service/internal/processor"
	repo "test-service/internal/repo/message"
	def "test-service/internal/service"
)

var _ def.MessageService = (*service)(nil)

type service struct {
	messageRepo      repo.MessageRepository
	messageProcessor processor.MessageProcessor
}

func NewService(msgRepo repo.MessageRepository, msgProcessor processor.MessageProcessor, cfg *config.Config) *service {
	return &service{
		messageRepo: msgRepo,
		messageProcessor:msgProcessor,
	}
}
