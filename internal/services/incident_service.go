package services

import (
	"JedelKomek/internal/models"
	"JedelKomek/internal/repositories"
	"context"
)

type IncidentService struct {
	Repo *repositories.IncidentRepository
}

func (s *IncidentService) Create(ctx context.Context, inc models.IncidentReport) (models.IncidentReport, error) {
	id, err := s.Repo.Create(ctx, inc)
	if err != nil {
		return models.IncidentReport{}, err
	}
	inc.ID = id
	return inc, nil
}

func (s *IncidentService) GetAll(ctx context.Context) ([]models.IncidentReport, error) {
	return s.Repo.GetAll(ctx)
}

func (s *IncidentService) GetByID(ctx context.Context, id int) (models.IncidentReport, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *IncidentService) Update(ctx context.Context, inc models.IncidentReport) (models.IncidentReport, error) {
	err := s.Repo.Update(ctx, inc)
	if err != nil {
		return models.IncidentReport{}, err
	}
	return s.Repo.GetByID(ctx, inc.ID)
}

func (s *IncidentService) Delete(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}
