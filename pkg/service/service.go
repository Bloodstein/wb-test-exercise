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
	GetOne(string) (*domain.TelegramToOfficeRelation, error)
	Create(row *domain.ModifyRequest) (string, error)
	Delete(string) (int, error)
	Update(string, *domain.ModifyRequest) (int, error)
	OfficeAlreadyExists(int) bool
	TelegramChatAlreadyExists(int) bool
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Relations: NewRelationsService(repo),
	}
}
