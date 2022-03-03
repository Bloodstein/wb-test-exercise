package main

import (
	"context"

	wbtestexercise "github.com/Bloodstein/wb-test-exercise"
	"github.com/Bloodstein/wb-test-exercise/pkg/handler"
	"github.com/Bloodstein/wb-test-exercise/pkg/repository"
	"github.com/Bloodstein/wb-test-exercise/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func loadConfigs() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("app")
	return viper.ReadInConfig()
}

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := loadConfigs(); err != nil {
		logrus.Fatalf("The reading configuration went wrong: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("The reading environment went wrong: %s", err.Error())
	}

	db, err := repository.NewMongoDB(&repository.Config{
		Host: viper.GetString("db.mongo.host"),
		Port: viper.GetString("db.mongo.port"),
	})

	if err != nil {
		logrus.Fatalf("Fail to connect to Mongo database: %s", err.Error())
	}

	err = db.Ping(context.Background(), &readpref.ReadPref{})

	if err != nil {
		logrus.Fatalf("Error to ping database: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	h := handler.NewHandler(service)

	if err := wbtestexercise.NewServer().Run(viper.GetString("port"), h.Routes()); err != nil {
		logrus.Fatalf("Error to run web server: %s", err.Error())
	}
}
