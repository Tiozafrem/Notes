package main

import (
	"notes"
	"notes/pkg/handler"
	"notes/pkg/repository"
	"notes/pkg/repository/postgres"
	"notes/pkg/usecases"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// @title Notes API
// @version 1.0
// @description Api server for Notes Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error load env config %s", err.Error())
	}

	if err := godotenv.Load("../.env"); err != nil {
		logrus.Fatalf("Error load env config %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db %s", err.Error())
	}

	repository := repository.NewRepository(db)
	usecases := usecases.NewUsecases(repository)
	handlers := handler.NewHandler(usecases)
	srv := new(notes.Server)
	srv.Run(viper.GetString("port"), handlers.InitRoutes())

}

// Init config from defaults pass
func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
