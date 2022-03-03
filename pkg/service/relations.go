package service

import (
	"errors"

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
	return r.repo.RelationsRepository.GetAll()
}

func (r *RelationsService) GetOne(rowId string) (*domain.TelegramToOfficeRelation, error) {
	return r.repo.RelationsRepository.GetOne(rowId)
}

func (r *RelationsService) Create(row *domain.ModifyRequest) (string, error) {

	if r.OfficeAlreadyExists(row.OfficeID) {
		return "", errors.New("That Office ID already exists")
	}

	if r.TelegramChatAlreadyExists(row.TelegramChatID) {
		return "", errors.New("That Telegram Chat ID already exists")
	}

	return r.repo.RelationsRepository.Create(row)
}

func (r *RelationsService) Delete(rowId string) (int, error) {
	return r.repo.RelationsRepository.Delete(rowId)
}

func (r *RelationsService) Update(rowId string, payload *domain.ModifyRequest) (int, error) {

	if r.OfficeAlreadyExists(payload.OfficeID) {
		return 0, errors.New("That Office ID already exists")
	}

	if r.TelegramChatAlreadyExists(payload.TelegramChatID) {
		return 0, errors.New("That Telegram Chat ID already exists")
	}

	return r.repo.RelationsRepository.Update(rowId, payload)
}
