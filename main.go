package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("520708190:AAHGJshPWGeDleVkmEPaCvUphErEvc5fTps")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Poshel nahui")
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}