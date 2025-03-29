package services

import (
	"JedelKomek/internal/models"
	"JedelKomek/internal/repositories"
	"context"
)

type MessageService struct {
	Repo *repositories.MessageRepository
}

func (s *MessageService) Create(ctx context.Context, msg models.Message) (models.Message, error) {
	id, err := s.Repo.Create(ctx, msg)
	if err != nil {
		return models.Message{}, err
	}
	msg.ID = id
	return msg, nil
}

func (s *MessageService) GetAll(ctx context.Context) ([]models.Message, error) {
	return s.Repo.GetAll(ctx)
}

func (s *MessageService) GetByID(ctx context.Context, id int) (models.Message, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *MessageService) Delete(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}
