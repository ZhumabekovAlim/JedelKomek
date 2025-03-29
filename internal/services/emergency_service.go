package services

import (
	"JedelKomek/internal/models"
	"JedelKomek/internal/repositories"
	"context"
)

type EmergencyService struct {
	Repo *repositories.EmergencyRepository
}

func (s *EmergencyService) Create(ctx context.Context, obj models.EmergencyCall) (models.EmergencyCall, error) {
	id, err := s.Repo.Create(ctx, obj)
	if err != nil {
		return models.EmergencyCall{}, err
	}
	obj.ID = id
	return obj, nil
}

func (s *EmergencyService) GetAll(ctx context.Context) ([]models.EmergencyCall, error) {
	return s.Repo.GetAll(ctx)
}

func (s *EmergencyService) GetByID(ctx context.Context, id int) (models.EmergencyCall, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *EmergencyService) Delete(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}
