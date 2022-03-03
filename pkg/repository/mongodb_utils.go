package repository

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *MongoDbRepository) collection() *mongo.Collection {
	return db.repo.Database(viper.GetString("db.mongo.database")).Collection(viper.GetString("db.mongo.collection"))
}

func (db *MongoDbRepository) OfficeAlreadyExists(officeId int) bool {

	result := db.collection().FindOne(ctx, bson.M{"officeid": bson.M{"$eq": officeId}})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return false
	}

	return result.Err() == nil
}

func (db *MongoDbRepository) TelegramChatAlreadyExists(telegramChatId int) bool {

	result := db.collection().FindOne(ctx, bson.M{"telegramchatid": bson.M{"$eq": telegramChatId}})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return false
	}

	return result.Err() == nil
}
