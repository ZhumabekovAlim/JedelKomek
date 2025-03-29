package services

import (
	"JedelKomek/internal/models"
	"JedelKomek/internal/repositories"
	"context"
)

type NewsService struct {
	Repo *repositories.NewsRepository
}

func (s *NewsService) Create(ctx context.Context, obj models.News) (models.News, error) {
	id, err := s.Repo.Create(ctx, obj)
	if err != nil {
		return models.News{}, err
	}
	obj.ID = id
	return obj, nil
}

func (s *NewsService) GetAll(ctx context.Context) ([]models.News, error) {
	return s.Repo.GetAll(ctx)
}

func (s *NewsService) GetByID(ctx context.Context, id int) (models.News, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *NewsService) Update(ctx context.Context, obj models.News) (models.News, error) {
	if err := s.Repo.Update(ctx, obj); err != nil {
		return models.News{}, err
	}
	return s.Repo.GetByID(ctx, obj.ID)
}

func (s *NewsService) Delete(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}
