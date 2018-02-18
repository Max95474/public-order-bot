package main

import (
	"public-order-bot/telegram"
	"public-order-bot/config"
)

func main() {
	telegram.Start(config.Config.TelegramKey, false)
}