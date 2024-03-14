package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	srv "github.com/takwot/tech-strelka.git"
	"github.com/takwot/tech-strelka.git/pkg/controllers"
	"github.com/takwot/tech-strelka.git/pkg/database"
	"github.com/takwot/tech-strelka.git/pkg/service"
)

func main(){

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("Failed in init config")
	}

	db, err := database.InitDb(&database.Config{
		Host:     "localhost",
		Username: "dev",
		SSLMode:  "disable",
		Password: "f7636c0azc8bfk1",
		DBName:   "techstrelka",
		Port:     "5555",
	})

	if err != nil {
		logrus.Fatalf("Error in init DB", err.Error())
	}

	repos := database.NewRepository(db)
	services := service.NewServices(repos)
	handlers := controllers.NewHandle(services)

	srv := new(srv.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error while start server!", err.Error())
	}

}

func initConfig() error{
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}