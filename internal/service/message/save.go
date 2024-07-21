package message

func (s *service) SaveMessage() error {
	err := s.messageRepo.SaveMessage()
	if err != nil {
		return err
	}
	return nil
}
