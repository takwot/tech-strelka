package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/takwot/tech-strelka.git/pkg/controllers"
	"github.com/takwot/tech-strelka.git/pkg/database"
	"github.com/takwot/tech-strelka.git/pkg/service"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("Failed in init config")
	}

	db, err := database.InitDb(&database.Config{
		Host:     "45.84.225.194",
		Username: "dev",
		SSLMode:  "disable",
		Password: "f7636c0azc8bfk1",
		DBName:   "techstrelka",
		Port:     "5555",
	})

	if err != nil {
		logrus.Fatal("Error in init DB", err.Error())
	}

	repos := database.NewRepository(db)
	services := service.NewServices(repos)
	handlers := controllers.NewHandle(services)

	srv := gin.Default()

	corsMiddleware := cors.AllowAll()

	srv.Use(corsMiddleware)
	handlers.InitRoutes(srv)

	srv.Run(":5000")
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
