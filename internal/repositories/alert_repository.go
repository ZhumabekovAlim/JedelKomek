package repositories

import (
	"JedelKomek/internal/models"
	"context"
	"database/sql"
)

type AlertRepository struct {
	Db *sql.DB
}

func (r *AlertRepository) Create(ctx context.Context, a models.Alert) (models.Alert, error) {
	query := `INSERT INTO alerts (title, description, status) VALUES (?, ?, ?)`
	res, err := r.Db.ExecContext(ctx, query, a.Title, a.Description, a.Status)
	if err != nil {
		return a, err
	}
	id, _ := res.LastInsertId()
	a.ID = int(id)
	return a, nil
}

func (r *AlertRepository) GetAll(ctx context.Context) ([]models.Alert, error) {
	query := `SELECT id, title, description, status, created_at FROM alerts ORDER BY created_at DESC`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Alert
	for rows.Next() {
		var a models.Alert
		if err := rows.Scan(&a.ID, &a.Title, &a.Description, &a.Status, &a.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

func (r *AlertRepository) GetByID(ctx context.Context, id int) (models.Alert, error) {
	query := `SELECT id, title, description, status, created_at FROM alerts WHERE id = ?`
	var a models.Alert
	err := r.Db.QueryRowContext(ctx, query, id).Scan(&a.ID, &a.Title, &a.Description, &a.Status, &a.CreatedAt)
	return a, err
}

func (r *AlertRepository) Update(ctx context.Context, a models.Alert) error {
	query := `UPDATE alerts SET title = ?, description = ?, status = ? WHERE id = ?`
	_, err := r.Db.ExecContext(ctx, query, a.Title, a.Description, a.Status, a.ID)
	return err
}

func (r *AlertRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM alerts WHERE id = ?`
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}
