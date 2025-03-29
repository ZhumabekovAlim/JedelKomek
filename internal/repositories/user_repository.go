package repositories

import (
	"JedelKomek/internal/models"
	"context"
	"database/sql"
)

type UserRepository struct {
	Db *sql.DB
}

func (r *UserRepository) Create(ctx context.Context, user models.User) (int, error) {
	query := `INSERT INTO users (fio, phone_number, password) VALUES (?, ?, ?)`
	res, err := r.Db.ExecContext(ctx, query, user.FIO, user.PhoneNumber, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	rows, err := r.Db.QueryContext(ctx, `SELECT id, fio, phone_number, password, created_at FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.ID, &u.FIO, &u.PhoneNumber, &u.Password, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (models.User, error) {
	var u models.User
	err := r.Db.QueryRowContext(ctx, `SELECT id, fio, phone_number, password, created_at FROM users WHERE id = ?`, id).
		Scan(&u.ID, &u.FIO, &u.PhoneNumber, &u.Password, &u.CreatedAt)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (r *UserRepository) Update(ctx context.Context, user models.User) error {
	query := `UPDATE users SET fio = ?, phone_number = ?, password = ? WHERE id = ?`
	_, err := r.Db.ExecContext(ctx, query, user.FIO, user.PhoneNumber, user.Password, user.ID)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, `DELETE FROM users WHERE id = ?`, id)
	return err
}
