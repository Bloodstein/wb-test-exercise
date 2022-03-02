package repository

import (
	"github.com/Bloodstein/wb-test-exercise/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	RelationsRepository
}

type RelationsRepository interface {
	GetAll() ([]*domain.TelegramToOfficeRelation, error)
	GetOne(int) (*domain.TelegramToOfficeRelation, error)
	Create(*domain.TelegramToOfficeRelation) (int, error)
	Update(int, *domain.TelegramToOfficeRelation) (int, error)
	Delete(int) (int, error)
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		RelationsRepository: NewMongoDbRepository(db),
	}
}
