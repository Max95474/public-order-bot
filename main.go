package main

import (
	"public-order-bot/telegram"
	"os"
)

func main() {
	 telegram.Start(os.Getenv("TELEGRAM_APIKEY"), false)
}
