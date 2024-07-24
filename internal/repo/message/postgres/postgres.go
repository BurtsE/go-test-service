package postgres

import (
	"database/sql"
	"test-service/internal/config"
	"test-service/internal/model"
	def "test-service/internal/repo/message"
)

var _ def.MessageRepository = (*repository)(nil)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB, cfg *config.Config) *repository {
	
	return &repository{
		db: db,
	}
}

func (r *repository) GetMessages() ([]model.Message, error) {
	query := `
		SELECT id, text, status
		FROM messages
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	messages := []model.Message{}
	for rows.Next() {
		msg := model.Message{}
		rows.Scan(&msg.Text, &msg.Status)
		messages = append(messages, msg)
	}
	return messages, nil
}

func (r *repository) SaveMessage(msg model.Message) error {
	query := `
		INSERT INTO messages (text, status)
		VALUES ($1, $2)
	`
	_, err := r.db.Query(query, msg.Text, msg.Status)
	if err != nil {
		return err
	}
	return nil
}
