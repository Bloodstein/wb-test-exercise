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
	GetOne(string) (*domain.TelegramToOfficeRelation, error)
	Create(*domain.ModifyRequest) (string, error)
	Update(string, *domain.ModifyRequest) (int, error)
	Delete(string) (int, error)
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		RelationsRepository: NewMongoDbRepository(db),
	}
}
