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

func (r *RelationsService) GetAll() ([]*domain.TelegramToOfficeRelation, error) {
	return r.repo.GetAll()
}

func (r *RelationsService) GetOne(rowId string) (*domain.TelegramToOfficeRelation, error) {
	return r.repo.GetOne(rowId)
}

func (r *RelationsService) Create(row *domain.ModifyRequest) (string, error) {
	return r.repo.Create(row)
}

func (r *RelationsService) Delete(rowId string) (int, error) {
	return r.repo.Delete(rowId)
}

func (r *RelationsService) Update(rowId string, payload *domain.ModifyRequest) (int, error) {
	return r.repo.Update(rowId, payload)
}
