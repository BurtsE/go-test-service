package app

import (
	"log"
	"test-service/internal/config"
	repo "test-service/internal/repo/message"
	"test-service/internal/repo/message/postgres"
	router "test-service/internal/router/message"
	"test-service/internal/service"
	"test-service/internal/service/message"

	"github.com/sirupsen/logrus"
)

type serviceProvider struct {
	cfg               *config.Config
	messageRepository repo.MessageRepository
	messageService    service.MessageService
	router            *router.Router
}

func NewSericeProvider() *serviceProvider {
	s := &serviceProvider{}
	s.Router()
	return s
}

func (s *serviceProvider) Config() *config.Config {
	if s.cfg == nil {
		cfg, err := config.InitConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.cfg = cfg
	}
	return s.cfg
}

func (s *serviceProvider) MessageRepository() repo.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = postgres.NewRepository(s.cfg)
	}
	return s.messageRepository
}

func (s *serviceProvider) MessageService() service.MessageService {
	if s.messageService == nil {
		s.messageService = message.NewService(s.MessageRepository(), s.cfg)
	}
	return s.messageService
}

func (s *serviceProvider) Router() *router.Router {
	if s.router == nil {
		s.router = router.NewRouter(logrus.New(), s.Config(), s.MessageService())
	}
	return s.router
}
