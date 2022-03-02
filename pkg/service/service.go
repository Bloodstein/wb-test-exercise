package service

import (
	"github.com/Bloodstein/wb-test-exercise/domain"
	"github.com/Bloodstein/wb-test-exercise/pkg/repository"
)

type Service struct {
	Relations
}

type Relations interface {
	GetAll() []*domain.TelegramToOfficeRelation
	GetOne(rowId int) *domain.TelegramToOfficeRelation
	Create(row *domain.TelegramToOfficeRelation) int
	Delete(rowId int) int
	Update(rowId int, payload *domain.TelegramToOfficeRelation) int
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Relations: NewRelationsService(repo),
	}
}
