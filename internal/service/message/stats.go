package message

import "test-service/internal/model"

func (s *service) GetStatistics() (uint, uint, error) {
	var unProcessed, processed uint

	messages, err := s.messageRepo.GetMessages()
	if err != nil {
		return 0, 0, err
	}

	for _, msg := range messages {
		switch msg.Status {
		case model.Finished:
			processed++
		case model.InProgress:
			unProcessed++
		}
	}
	return unProcessed, processed, nil
}
