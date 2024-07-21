package message

import "test-service/internal/model"

func (s *service) GetMessages() ([]model.Message, error) {
	messages, err := s.messageRepo.GetMessages()
	if err != nil {
		return nil, err
	}
	return messages, nil
}
