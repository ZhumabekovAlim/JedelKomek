package repositories

import (
	"JedelKomek/internal/models"
	"context"
	"database/sql"
)

type EmergencyRepository struct {
	Db *sql.DB
}

func (r *EmergencyRepository) Create(ctx context.Context, obj models.EmergencyCall) (int, error) {
	query := `INSERT INTO emergency_calls (user_id, latitude, longitude) VALUES (?, ?, ?)`
	res, err := r.Db.ExecContext(ctx, query, obj.UserID, obj.Latitude, obj.Longitude)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (r *EmergencyRepository) GetAll(ctx context.Context) ([]models.EmergencyCall, error) {
	rows, err := r.Db.QueryContext(ctx, `SELECT id, user_id, latitude, longitude, created_at FROM emergency_calls`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.EmergencyCall
	for rows.Next() {
		var obj models.EmergencyCall
		err := rows.Scan(&obj.ID, &obj.UserID, &obj.Latitude, &obj.Longitude, &obj.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

func (r *EmergencyRepository) GetByID(ctx context.Context, id int) (models.EmergencyCall, error) {
	var obj models.EmergencyCall
	query := `SELECT id, user_id, latitude, longitude, created_at FROM emergency_calls WHERE id = ?`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(&obj.ID, &obj.UserID, &obj.Latitude, &obj.Longitude, &obj.CreatedAt)
	return obj, err
}

func (r *EmergencyRepository) Delete(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, `DELETE FROM emergency_calls WHERE id = ?`, id)
	return err
}
