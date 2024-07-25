package postgres

import (
	"database/sql"
	"fmt"
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
		err := rows.Scan(&msg.UUID, &msg.Text, &msg.Status)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func (r *repository) SaveMessage(msg model.Message) (uint64, error) {
	var id uint64
	query := `
		INSERT INTO messages (text, status)
		VALUES ($1, $2)
		RETURNING id
	`
	row := r.db.QueryRow(query, msg.Text, msg.Status)
	if row.Err() != nil {
		return 0, row.Err()
	}
	row.Scan(&id)
	return id, nil
}

func (r *repository) UpdateMessageStatus(msg model.Message) error {
	query := `
		UPDATE messages
		SET status = $2
		WHERE id = $1
	`
	_, err := r.db.Query(query, msg.UUID, msg.Status)
	if err != nil {
		return err
	}
	return nil
}
