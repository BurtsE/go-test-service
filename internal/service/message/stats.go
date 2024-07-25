package message

import (
	"test-service/internal/model"
)

func (s *service) GetStatistics() (model.Statistics, error) {
	var unProcessed, processed, received uint
	stats := model.Statistics{}
	messages, err := s.messageRepo.GetMessages()
	if err != nil {
		return stats, err
	}
	for _, msg := range messages {
		switch msg.Status {
		case model.Finished:
			processed++
		case model.InProgress:
			unProcessed++
		case model.Created:
			received++
		}
	}
	stats.InProgress = &unProcessed
	stats.Processed = &processed
	stats.Received = &received
	return stats, nil
}
