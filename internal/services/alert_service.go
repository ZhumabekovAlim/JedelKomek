package services

import (
	"JedelKomek/internal/models"
	"JedelKomek/internal/repositories"
	"context"
)

type AlertService struct {
	Repo *repositories.AlertRepository
}

func (s *AlertService) Create(ctx context.Context, a models.Alert) (models.Alert, error) {
	return s.Repo.Create(ctx, a)
}

func (s *AlertService) GetAll(ctx context.Context) ([]models.Alert, error) {
	return s.Repo.GetAll(ctx)
}

func (s *AlertService) GetByID(ctx context.Context, id int) (models.Alert, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *AlertService) Update(ctx context.Context, a models.Alert) error {
	return s.Repo.Update(ctx, a)
}

func (s *AlertService) Delete(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}
