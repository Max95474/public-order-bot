package main

import (
	"public-order-bot/telegram"
)

func main() {
	telegram.Start("520708190:AAHGJshPWGeDleVkmEPaCvUphErEvc5fTps", false)
	// telegram.Start(os.Getenv("TELEGRAM_APIKEY"))
}