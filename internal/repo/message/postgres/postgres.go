package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"test-service/internal/config"
	"test-service/internal/model"
	def "test-service/internal/repo/message"

	_ "github.com/lib/pq"
)

var _ def.MessageRepository = (*repository)(nil)

type repository struct {
	db *sql.DB
}

func NewRepository(cfg *config.Config) (*repository, error) {
	log.Println("initiating repository")
	DSN := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		cfg.DB,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Postgres.Port,
		cfg.Sslmode,
	)
	db, _ := sql.Open("postgres", DSN)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &repository{
		db: db,
	}, nil
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
