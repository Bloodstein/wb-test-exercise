package repository

import "github.com/Bloodstein/wb-test-exercise/domain"

type Repository struct {
	RelationsRepository
}

type RelationsRepository interface {
	GetAll() []*domain.TelegramToOfficeRelation
	GetOne() *domain.TelegramToOfficeRelation
	Create() int
	Remove() int
	Update() int
	Delete() int
}

func NewRepository() *Repository {
	return &Repository{
		RelationsRepository: NewMongoDbRepository(),
	}
}
