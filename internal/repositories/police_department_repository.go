package repositories

import (
	"JedelKomek/internal/models"
	"context"
	"database/sql"
)

type PoliceDepartmentRepository struct {
	Db *sql.DB
}

func (r *PoliceDepartmentRepository) Create(ctx context.Context, pd models.PoliceDepartment) (int, error) {
	query := `INSERT INTO police_department (name, phone_number, address, work_days, work_time, latitude, longitude) VALUES (?, ?, ?, ?, ?, ?, ?)`
	res, err := r.Db.ExecContext(ctx, query, pd.Name, pd.PhoneNumber, pd.Address, pd.WorkDays, pd.WorkTime, pd.Latitude, pd.Longitude)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (r *PoliceDepartmentRepository) GetAll(ctx context.Context) ([]models.PoliceDepartment, error) {
	rows, err := r.Db.QueryContext(ctx, `SELECT id, name, phone_number, address, work_days, work_time, latitude, longitude, created_at FROM police_department`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.PoliceDepartment
	for rows.Next() {
		var pd models.PoliceDepartment
		err := rows.Scan(&pd.ID, &pd.Name, &pd.PhoneNumber, &pd.Address, &pd.WorkDays, &pd.WorkTime, &pd.Latitude, &pd.Longitude, &pd.CreatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, pd)
	}
	return list, nil
}

func (r *PoliceDepartmentRepository) GetByID(ctx context.Context, id int) (models.PoliceDepartment, error) {
	var pd models.PoliceDepartment
	query := `SELECT id, name, phone_number, address, work_days, work_time, latitude, longitude, created_at FROM police_department WHERE id = ?`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(&pd.ID, &pd.Name, &pd.PhoneNumber, &pd.Address, &pd.WorkDays, &pd.WorkTime, &pd.Latitude, &pd.Longitude, &pd.CreatedAt)
	return pd, err
}

func (r *PoliceDepartmentRepository) Update(ctx context.Context, pd models.PoliceDepartment) error {
	query := `UPDATE police_department SET name = ?, phone_number = ?, address = ?, work_days = ?, work_time = ?, latitude = ?, longitude = ? WHERE id = ?`
	_, err := r.Db.ExecContext(ctx, query, pd.Name, pd.PhoneNumber, pd.Address, pd.WorkDays, pd.WorkTime, pd.Latitude, pd.Longitude, pd.ID)
	return err
}

func (r *PoliceDepartmentRepository) Delete(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, `DELETE FROM police_department WHERE id = ?`, id)
	return err
}
