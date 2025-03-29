package repositories

import (
	"JedelKomek/internal/models"
	"context"
	"database/sql"
)

type EducationRepository struct {
	Db *sql.DB
}

func (r *EducationRepository) Create(ctx context.Context, obj models.EducationContent) (int, error) {
	query := `INSERT INTO education_contents (title, body, media_url) VALUES (?, ?, ?)`
	res, err := r.Db.ExecContext(ctx, query, obj.Title, obj.Body, obj.MediaURL)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (r *EducationRepository) GetAll(ctx context.Context) ([]models.EducationContent, error) {
	rows, err := r.Db.QueryContext(ctx, `SELECT id, title, body, media_url, created_at FROM education_contents`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.EducationContent
	for rows.Next() {
		var obj models.EducationContent
		err = rows.Scan(&obj.ID, &obj.Title, &obj.Body, &obj.MediaURL, &obj.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

func (r *EducationRepository) GetByID(ctx context.Context, id int) (models.EducationContent, error) {
	var obj models.EducationContent
	err := r.Db.QueryRowContext(ctx, `SELECT id, title, body, media_url, created_at FROM education_contents WHERE id = ?`, id).
		Scan(&obj.ID, &obj.Title, &obj.Body, &obj.MediaURL, &obj.CreatedAt)
	if err != nil {
		return models.EducationContent{}, err
	}
	return obj, nil
}

func (r *EducationRepository) Update(ctx context.Context, obj models.EducationContent) error {
	query := `UPDATE education_contents SET title = ?, body = ?, media_url = ? WHERE id = ?`
	_, err := r.Db.ExecContext(ctx, query, obj.Title, obj.Body, obj.MediaURL, obj.ID)
	return err
}

func (r *EducationRepository) Delete(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, `DELETE FROM education_contents WHERE id = ?`, id)
	return err
}
