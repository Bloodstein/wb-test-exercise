package repository

import (
	"context"
	"fmt"

	"github.com/Bloodstein/wb-test-exercise/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const connection_string = "mongodb+srv://saomlydb:Ghbdtn_123@cluster0.dqzx8.mongodb.net/Cluster0?retryWrites=true&w=majority"
const (
	connection_string = "mongodb://saomlydb:Ghbdtn_123@cluster0.dqzx8.mongodb.net/Cluster0?retryWrites=true&w=majority"
)

var ctx = context.TODO()

type Config struct {
	Database string
	Login    string
	Password string
}

type MongoDbRepository struct {
	repo *mongo.Client
}

func NewMongoDB(conf *Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%s:%s@cluster0.dqzx8.mongodb.net/%s?retryWrites=true&w=majority", conf.Login, conf.Password, conf.Database),
	)

	return mongo.Connect(ctx, clientOptions)
}

func NewMongoDbRepository(db *mongo.Client) *MongoDbRepository {
	return &MongoDbRepository{
		repo: db,
	}
}

func (db *MongoDbRepository) collection() *mongo.Collection {
	return db.repo.Database(viper.GetString("db.mongo.database")).Collection(viper.GetString("db.mongo.collection"))
}

func (db *MongoDbRepository) GetAll() ([]*domain.TelegramToOfficeRelation, error) {
	collection := db.collection()

	var rows []*domain.TelegramToOfficeRelation

	cur, err := collection.Find(ctx, bson.D{{}})

	if err != nil {
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

func (db *MongoDbRepository) GetOne(rowId int) (*domain.TelegramToOfficeRelation, error) {

	result := db.collection().FindOne(ctx, bson.M{"id": rowId})

	if result.Err() == mongo.ErrNoDocuments {
		return nil, result.Err()
	}

	var data domain.TelegramToOfficeRelation
	err := result.Decode(&data)

	return &data, err
}

func (db *MongoDbRepository) Create(input *domain.TelegramToOfficeRelation) (int, error) {

	result, err := db.collection().InsertOne(ctx, &input)

	if err != nil {
		return 0, err
	}

	return result.InsertedID.(int), nil
}

func (db *MongoDbRepository) Update(rowId int, input *domain.TelegramToOfficeRelation) (int, error) {

	_, err := db.collection().UpdateOne(ctx, bson.M{"id": rowId}, &input)

	if err != nil {
		return 0, err
	}

	return rowId, nil
}

func (db *MongoDbRepository) Delete(rowId int) (int, error) {

	_, err := db.collection().DeleteOne(ctx, bson.M{"id": rowId})

	if err != nil {
		return 0, err
	}

	return rowId, nil
}
