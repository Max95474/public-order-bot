package telegram

import (
  "github.com/Syfaro/telegram-bot-api"
  "fmt"
  "io/ioutil"
  "text/template"
  "bytes"
)

func getResponse(update *tgbotapi.Update) {
  switch update.Message.Command() {
    case "create": createOrder(update)
    case "list": ordersList(update)
    case "add": addItem(update)
    case "remove": removeItem(update)
    case "result": orderResult(update)
    case "close": closeOrder(update)
  }
}

func createOrder(update *tgbotapi.Update) {
  orderName := update.Message.CommandArguments()
  var response string
  if orderName == "" {
    response = "Enter order name i.e `/create dinner indian food`"
  } else {
    dat, err := ioutil.ReadFile("./telegram/templates/create_success.md")
    check(err)
    tmpl := template.New("test")
    tmpl, err = tmpl.Parse(string("Your order *{{.orderName}}* has been created"))
    check(err)
    response = string(dat)
    fmt.Println(response)
  }
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func ordersList(update *tgbotapi.Update) {
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, "azazaza")
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func addItem(update *tgbotapi.Update) {
  response := "Item added to order"
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func removeItem(update *tgbotapi.Update) {
  response := "Item has been removed"
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func orderResult(update *tgbotapi.Update) {
  response := "Order result"
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func closeOrder(update *tgbotapi.Update) {
  response := "Order has been closed"
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}
