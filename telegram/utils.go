package telegram

import "github.com/Syfaro/telegram-bot-api"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func sendMessage(chatId int64, msgData string) {
  msg := tgbotapi.NewMessage(chatId, msgData)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}
