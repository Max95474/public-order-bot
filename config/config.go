package config

import (
	"os"
)

type config struct {
	TelegramKey string
	Database struct {
	  Host         string
		Username     string
		Password     string
		DatabaseName string
	}
}

var Config *config

func init() {
	Config = loadConfig()
}

func loadConfig() *config {
	appConfig := &config{}

	appConfig.TelegramKey = os.Getenv("TELEGRAM_APIKEY")
	appConfig.Database.Host = os.Getenv("POSTGRES_HOST")
	appConfig.Database.Username = os.Getenv("POSTGRES_USER")
	appConfig.Database.Password = os.Getenv("POSTGRES_PASSWORD")
	appConfig.Database.DatabaseName = os.Getenv("POSTGRES_DB")

	return appConfig
}
