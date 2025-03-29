package services

import (
	"JedelKomek/internal/models"
	"JedelKomek/internal/repositories"
	"context"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) Create(ctx context.Context, user models.User) (models.User, error) {
	id, err := s.Repo.Create(ctx, user)
	if err != nil {
		return models.User{}, err
	}
	user.ID = id
	return user, nil
}

func (s *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.Repo.GetAll(ctx)
}

func (s *UserService) GetByID(ctx context.Context, id int) (models.User, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, user models.User) (models.User, error) {
	err := s.Repo.Update(ctx, user)
	if err != nil {
		return models.User{}, err
	}
	return s.Repo.GetByID(ctx, user.ID)
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}
