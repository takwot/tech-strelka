package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/takwot/tech-strelka.git/pkg/controllers"
	"github.com/takwot/tech-strelka.git/pkg/database"
	"github.com/takwot/tech-strelka.git/pkg/service"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	cfg := Config{}
	if err := initConfig(&cfg); err != nil {
		logrus.Fatalf("Failed in init config")
	}

	db, err := database.InitDb(cfg.Postgres)
	if err != nil {
		logrus.Fatal("Error in init DB", err.Error())
	}

	repos := database.NewRepository(db)
	services := service.NewServices(repos)
	handlers := controllers.NewHandle(services)

	srv := gin.Default()

	handlers.InitRoutes(srv)

	if err := srv.Run(cfg.Server.Addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatal(fmt.Sprintf("problem with run server: %+v", err))
	}
}

func initConfig(cfg *Config) error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}
	err = viper.Unmarshal(cfg)

	return err
}
