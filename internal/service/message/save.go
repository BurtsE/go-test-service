package message

import "test-service/internal/model"

func (s *service) SaveMessage(msg model.Message) error {
	msg.Status = model.Created
	err := s.messageRepo.SaveMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
