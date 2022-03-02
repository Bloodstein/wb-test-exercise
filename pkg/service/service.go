package service

import (
	"github.com/Bloodstein/wb-test-exercise/domain"
	"github.com/Bloodstein/wb-test-exercise/pkg/repository"
)

type Service struct {
	Relations
}

type Relations interface {
	GetAll() ([]*domain.TelegramToOfficeRelation, error)
	GetOne(int) (*domain.TelegramToOfficeRelation, error)
	Create(row *domain.TelegramToOfficeRelation) (int, error)
	Delete(int) (int, error)
	Update(int, *domain.TelegramToOfficeRelation) (int, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Relations: NewRelationsService(repo),
	}
}
