package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatal("Error load env")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load env")
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
