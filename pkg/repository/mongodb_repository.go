package repository

import "github.com/Bloodstein/wb-test-exercise/domain"

type MongoDbRepository struct{}

func NewMongoDbRepository() *MongoDbRepository {
	return &MongoDbRepository{}
}

func (db *MongoDbRepository) GetAll() []*domain.TelegramToOfficeRelation {

}

func (db *MongoDbRepository) GetOne() *domain.TelegramToOfficeRelation {

}

func (db *MongoDbRepository) Create() int {

}

func (db *MongoDbRepository) Remove() int {

}

func (db *MongoDbRepository) Update() int {

}

func (db *MongoDbRepository) Delete() int {

}
