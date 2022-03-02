package main

import (
	wbtestexercise "github.com/Bloodstein/wb-test-exercise"
	"github.com/Bloodstein/wb-test-exercise/pkg/handler"
	"github.com/Bloodstein/wb-test-exercise/pkg/repository"
	"github.com/Bloodstein/wb-test-exercise/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func loadConfigs() error {
	viper.AddConfigPath("config")
	viper.SetConfigFile("app")
	return viper.ReadInConfig()
}

func main() {

	if err := loadConfigs(); err != nil {
		logrus.Fatalf("The reading configuration went wrong: %s", err)
	}

	repo := repository.NewRepository()
	service := service.NewService(repo)
	h := handler.NewHandler(service)

	server := wbtestexercise.NewServer()
	server.Run(viper.GetString("port"), h.Routes())
}
