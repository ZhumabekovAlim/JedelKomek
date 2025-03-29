package repositories

import (
	"JedelKomek/internal/models"
	"context"
	"database/sql"
)

type MessageRepository struct {
	Db *sql.DB
}

func (r *MessageRepository) Create(ctx context.Context, msg models.Message) (int, error) {
	query := `INSERT INTO messages (sender_id, receiver_id, text) VALUES (?, ?, ?)`
	res, err := r.Db.ExecContext(ctx, query, msg.SenderID, msg.ReceiverID, msg.Text)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (r *MessageRepository) GetAll(ctx context.Context) ([]models.Message, error) {
	rows, err := r.Db.QueryContext(ctx, `SELECT id, sender_id, receiver_id, text, created_at FROM messages`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err = rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Text, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func (r *MessageRepository) GetByID(ctx context.Context, id int) (models.Message, error) {
	var msg models.Message
	query := `SELECT id, sender_id, receiver_id, text, created_at FROM messages WHERE id = ?`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Text, &msg.CreatedAt)
	return msg, err
}

func (r *MessageRepository) Delete(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, `DELETE FROM messages WHERE id = ?`, id)
	return err
}
