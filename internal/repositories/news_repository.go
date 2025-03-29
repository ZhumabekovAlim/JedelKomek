package repositories

import (
	"JedelKomek/internal/models"
	"context"
	"database/sql"
)

type NewsRepository struct {
	Db *sql.DB
}

func (r *NewsRepository) Create(ctx context.Context, obj models.News) (int, error) {
	query := `INSERT INTO news (title, content, media_url) VALUES (?, ?, ?)`
	res, err := r.Db.ExecContext(ctx, query, obj.Title, obj.Content, obj.MediaURL)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (r *NewsRepository) GetAll(ctx context.Context) ([]models.News, error) {
	rows, err := r.Db.QueryContext(ctx, `SELECT id, title, content, media_url, created_at, updated_at FROM news`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.News
	for rows.Next() {
		var obj models.News
		err = rows.Scan(&obj.ID, &obj.Title, &obj.Content, &obj.MediaURL, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

func (r *NewsRepository) GetByID(ctx context.Context, id int) (models.News, error) {
	var obj models.News
	query := `SELECT id, title, content, media_url, created_at, updated_at FROM news WHERE id = ?`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(&obj.ID, &obj.Title, &obj.Content, &obj.MediaURL, &obj.CreatedAt, &obj.UpdatedAt)
	return obj, err
}

func (r *NewsRepository) Update(ctx context.Context, obj models.News) error {
	query := `UPDATE news SET title = ?, content = ?, media_url = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	_, err := r.Db.ExecContext(ctx, query, obj.Title, obj.Content, obj.MediaURL, obj.ID)
	return err
}

func (r *NewsRepository) Delete(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, `DELETE FROM news WHERE id = ?`, id)
	return err
}
