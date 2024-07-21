package postgres

import (
	"database/sql"
	"test-service/internal/model"
	def "test-service/internal/service"
)

var _ def.MessageService = (*repository)(nil)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (s *repository) GetMessages() ([]model.Message, error) {
	return nil, nil
}

func (s *repository) SaveMessage() error {
	return nil
}
