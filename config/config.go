package config

import (
	"os"
)

type config struct {
	TelegramKey string
	Database struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		DatabaseName string `json:"database_name"`
	} `json:"database"`
}

var Config *config

func init() {
	Config = loadConfig()
}

func loadConfig() *config {
	appConfig := &config{}

	appConfig.TelegramKey = os.Getenv("TELEGRAM_APIKEY")
	appConfig.Database.Username = os.Getenv("POSTGRES_USER")
	appConfig.Database.Password = os.Getenv("POSTGRES_PASSWORD")
	appConfig.Database.DatabaseName = os.Getenv("POSTGRES_DB")

	return appConfig
}