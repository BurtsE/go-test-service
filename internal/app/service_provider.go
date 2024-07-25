package app

import (
	"log"
	"test-service/internal/config"
	"test-service/internal/processor"
	kafka "test-service/internal/processor/message"
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
	messageProcessor  processor.MessageProcessor
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
		repo, err := postgres.NewRepository(s.cfg)
		if err != nil {
			log.Fatalf("could not init repo: %s", err.Error())
		}
		s.messageRepository = repo
	}
	return s.messageRepository
}

func (s *serviceProvider) MessageProcessor() processor.MessageProcessor {
	if s.messageProcessor == nil {
		p, err := kafka.NewKafka()
		if err != nil {
			log.Fatal("could not connect to kafka:", err)
		}
		s.messageProcessor = p
	}
	return s.messageProcessor
}

func (s *serviceProvider) MessageService() service.MessageService {
	if s.messageService == nil {
		s.messageService = message.NewService(s.MessageRepository(), s.MessageProcessor(), s.cfg)
	}
	return s.messageService
}

func (s *serviceProvider) Router() *router.Router {
	if s.router == nil {
		s.router = router.NewRouter(logrus.New(), s.Config(), s.MessageService())
	}
	return s.router
}
