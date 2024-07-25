package message

import "test-service/internal/model"

func (s *service) SaveMessage(msg model.Message) error {
	msg.Status = model.Created
	id, err := s.messageRepo.SaveMessage(msg)
	if err != nil {
		return err
	}
	
	msg.UUID = id
	err = s.messageProcessor.SendMessage(msg)
	if err != nil {
		return err
	}

	msg.Status = model.InProgress
	err = s.messageRepo.UpdateMessageStatus(msg)
	if err != nil {
		return err
	}
	return nil
}
