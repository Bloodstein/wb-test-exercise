package repository

import (
	"context"
	"fmt"

	"github.com/Bloodstein/wb-test-exercise/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type Config struct {
	Host string
	Port string
}

type MongoDbRepository struct {
	repo *mongo.Client
}

func NewMongoDB(conf *Config) (*mongo.Client, error) {

	connString := fmt.Sprintf("mongodb://%s:%s", conf.Host, conf.Port)

	clientOptions := options.Client().ApplyURI(
		connString,
	)

	return mongo.Connect(ctx, clientOptions)
}

func NewMongoDbRepository(db *mongo.Client) *MongoDbRepository {
	return &MongoDbRepository{
		repo: db,
	}
}

// Получить все записи
func (db *MongoDbRepository) GetAll() ([]*domain.TelegramToOfficeRelation, error) {
	collection := db.collection()

	var rows []*domain.TelegramToOfficeRelation

	cur, err := collection.Find(ctx, bson.D{{}})

	if cur.RemainingBatchLength() == 0 {
		return rows, nil
	}

	if err != nil && err != mongo.ErrNoDocuments {
		return rows, err
	}

	for cur.Next(ctx) {
		var t domain.TelegramToOfficeRelation
		err := cur.Decode(&t)
		if err != nil {
			return rows, err
		}

		rows = append(rows, &t)
	}

	if err := cur.Err(); err != nil {
		return rows, err
	}

	cur.Close(ctx)

	if len(rows) == 0 {
		return rows, mongo.ErrNoDocuments
	}

	return rows, nil
}

// Получить одну конкретную запись по _id
func (db *MongoDbRepository) GetOne(rowId string) (*domain.TelegramToOfficeRelation, error) {

	objectId, err := primitive.ObjectIDFromHex(rowId)

	if err != nil {
		return nil, err
	}

	result := db.collection().FindOne(ctx, bson.M{"_id": bson.M{"$eq": objectId}})

	if result.Err() != nil {
		return nil, result.Err()
	}

	var data domain.TelegramToOfficeRelation

	err = result.Decode(&data)

	return &data, err
}

// Создать новую запись в БД
func (db *MongoDbRepository) Create(input *domain.ModifyRequest) (string, error) {

	result, err := db.collection().InsertOne(ctx, &input)

	if err != nil {
		return "", err
	}

	newId := result.InsertedID.(primitive.ObjectID)

	return newId.Hex(), nil
}

// Изменить существующую запись по _id
func (db *MongoDbRepository) Update(rowId string, input *domain.ModifyRequest) (int, error) {

	objectId, _ := primitive.ObjectIDFromHex(rowId)
	result, err := db.collection().UpdateOne(ctx, bson.M{"_id": bson.M{"$eq": objectId}}, bson.M{"$set": &input})

	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

// Удалить существующую запись по _id
func (db *MongoDbRepository) Delete(rowId string) (int, error) {

	objectId, _ := primitive.ObjectIDFromHex(rowId)
	result, err := db.collection().DeleteOne(ctx, bson.M{"_id": bson.M{"$eq": objectId}})

	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}
