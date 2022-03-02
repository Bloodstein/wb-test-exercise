package service

import (
	"github.com/Bloodstein/wb-test-exercise/domain"
	"github.com/Bloodstein/wb-test-exercise/pkg/repository"
)

type RelationsService struct {
	repo *repository.Repository
}

func NewRelationsService(repo *repository.Repository) *RelationsService {
	return &RelationsService{repo: repo}
}

func (r *RelationsService) GetAll() []*domain.TelegramToOfficeRelation {

}

func (r *RelationsService) GetOne(rowId int) *domain.TelegramToOfficeRelation {

}

func (r *RelationsService) Create(row *domain.TelegramToOfficeRelation) int {

}

func (r *RelationsService) Delete(rowId int) int {

}

func (r *RelationsService) Update(rowId int, payload *domain.TelegramToOfficeRelation) int {

}
