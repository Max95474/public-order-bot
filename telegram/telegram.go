package telegram

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)


func Start(apiKey string, debug bool) {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	
  if err != nil {
		log.Panic(err)
	}

	bot.Debug = debug;

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		
    command := ejectCommand(update.Message.Text, bot.Self.UserName)

		response := getResponse(command)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}