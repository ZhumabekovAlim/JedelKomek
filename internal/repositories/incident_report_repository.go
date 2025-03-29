package repositories

import (
	"JedelKomek/internal/models"
	"context"
	"database/sql"
)

type IncidentRepository struct {
	Db *sql.DB
}

func (r *IncidentRepository) Create(ctx context.Context, inc models.IncidentReport) (int, error) {
	query := `INSERT INTO incident_reports (user_id, description, media_url, type_id, latitude, longitude) 
			  VALUES (?, ?, ?, ?, ?, ?)`
	result, err := r.Db.ExecContext(ctx, query, inc.UserID, inc.Description, inc.MediaURL, inc.TypeID, inc.Latitude, inc.Longitude)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (r *IncidentRepository) GetAll(ctx context.Context) ([]models.IncidentReport, error) {
	query := `SELECT id, user_id, description, media_url, type_id, latitude, longitude, created_at FROM incident_reports`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incidents []models.IncidentReport
	for rows.Next() {
		var inc models.IncidentReport
		if err := rows.Scan(&inc.ID, &inc.UserID, &inc.Description, &inc.MediaURL, &inc.TypeID, &inc.Latitude, &inc.Longitude, &inc.CreatedAt); err != nil {
			return nil, err
		}
		incidents = append(incidents, inc)
	}
	return incidents, nil
}

func (r *IncidentRepository) GetByID(ctx context.Context, id int) (models.IncidentReport, error) {
	query := `SELECT id, user_id, description, media_url, type_id, latitude, longitude, created_at 
			  FROM incident_reports WHERE id = ?`
	var inc models.IncidentReport
	err := r.Db.QueryRowContext(ctx, query, id).Scan(&inc.ID, &inc.UserID, &inc.Description, &inc.MediaURL, &inc.TypeID, &inc.Latitude, &inc.Longitude, &inc.CreatedAt)
	if err != nil {
		return models.IncidentReport{}, err
	}
	return inc, nil
}

func (r *IncidentRepository) Update(ctx context.Context, inc models.IncidentReport) error {
	query := `UPDATE incident_reports 
			  SET description = ?, media_url = ?, type_id = ?, latitude = ?, longitude = ? 
			  WHERE id = ?`
	_, err := r.Db.ExecContext(ctx, query, inc.Description, inc.MediaURL, inc.TypeID, inc.Latitude, inc.Longitude, inc.ID)
	return err
}

func (r *IncidentRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM incident_reports WHERE id = ?`
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}
