package message

import (
	"database/sql"
	"log/slog"
)

type MessageRepository struct {
	db *sql.DB

	logger *slog.Logger
}

func NewRepo(db *sql.DB, logger *slog.Logger) *MessageRepository {
	return &MessageRepository{
		db:     db,
		logger: logger,
	}
}

func (r *MessageRepository) SaveMsg(msg string) error {
	query := "INSERT INTO messages (content) VALUES (?)"
	_, err := r.db.Exec(query, msg)

	return err
}
