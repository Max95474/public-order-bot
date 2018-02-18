package telegram

import (
  "github.com/Syfaro/telegram-bot-api"
  "log"
)

var telegramBot *tgbotapi.BotAPI

func Start(apiKey string, debug bool) {
  bot, err := tgbotapi.NewBotAPI(apiKey)
  telegramBot = bot
  if err != nil {
    log.Panic(err)
  }
  telegramBot.Debug = debug
  log.Printf("Authorized on account %s", telegramBot.Self.UserName)
  u := tgbotapi.NewUpdate(0)
  u.Timeout = 60
  updates, err := telegramBot.GetUpdatesChan(u)
  for update := range updates {
    if update.Message == nil {
      continue
    }
    log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
    getResponse(&update)
  }
}
