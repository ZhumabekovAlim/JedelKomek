package services

import (
	"JedelKomek/internal/models"
	"JedelKomek/internal/repositories"
	"context"
)

type EducationService struct {
	Repo *repositories.EducationRepository
}

func (s *EducationService) Create(ctx context.Context, obj models.EducationContent) (models.EducationContent, error) {
	id, err := s.Repo.Create(ctx, obj)
	if err != nil {
		return models.EducationContent{}, err
	}
	obj.ID = id
	return obj, nil
}

func (s *EducationService) GetAll(ctx context.Context) ([]models.EducationContent, error) {
	return s.Repo.GetAll(ctx)
}

func (s *EducationService) GetByID(ctx context.Context, id int) (models.EducationContent, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *EducationService) Update(ctx context.Context, obj models.EducationContent) (models.EducationContent, error) {
	err := s.Repo.Update(ctx, obj)
	if err != nil {
		return models.EducationContent{}, err
	}
	return s.Repo.GetByID(ctx, obj.ID)
}

func (s *EducationService) Delete(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}
