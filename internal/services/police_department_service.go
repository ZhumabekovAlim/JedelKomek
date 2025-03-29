package services

import (
	"JedelKomek/internal/models"
	"JedelKomek/internal/repositories"
	"context"
)

type PoliceDepartmentService struct {
	Repo *repositories.PoliceDepartmentRepository
}

func (s *PoliceDepartmentService) Create(ctx context.Context, pd models.PoliceDepartment) (models.PoliceDepartment, error) {
	id, err := s.Repo.Create(ctx, pd)
	if err != nil {
		return models.PoliceDepartment{}, err
	}
	pd.ID = id
	return pd, nil
}

func (s *PoliceDepartmentService) GetAll(ctx context.Context) ([]models.PoliceDepartment, error) {
	return s.Repo.GetAll(ctx)
}

func (s *PoliceDepartmentService) GetByID(ctx context.Context, id int) (models.PoliceDepartment, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *PoliceDepartmentService) Update(ctx context.Context, pd models.PoliceDepartment) (models.PoliceDepartment, error) {
	err := s.Repo.Update(ctx, pd)
	if err != nil {
		return models.PoliceDepartment{}, err
	}
	return s.Repo.GetByID(ctx, pd.ID)
}

func (s *PoliceDepartmentService) Delete(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}
