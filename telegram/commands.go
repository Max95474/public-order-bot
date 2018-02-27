package telegram

import (
  "github.com/Syfaro/telegram-bot-api"
  "public-order-bot/dao"
  "strconv"
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
    response = getResponseText("create_fail_order_name", nil)
  } else {
    order, err := dao.CreateOrder(string(update.Message.Chat.ID), string(update.Message.From.UserName), orderName)
    check(err)
    responseData := templateData {
      "orderName": order.Name,
      "orderId": order.Id,
    }
    response = getResponseText("create_success", responseData)
  }
  msg := tgbotapi.NewMessage(int64(update.Message.From.ID), response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func ordersList(update *tgbotapi.Update) {
  orders, err := dao.ListOrders(string(update.Message.Chat.ID))
  check(err)
  responseData := templateData {
    "orderList": orders,
  }
  response := getResponseText("orders_list", responseData)
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  telegramBot.Send(msg)
}

func addItem(update *tgbotapi.Update) {
  response := "Item added to order"
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func removeItem(update *tgbotapi.Update) {
  err := dao.DeleteOrder(string(update.Message.Chat.ID), string(update.Message.From.ID), 12)
  check(err)
  if err == nil {

  }
  var response string
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func orderResult(update *tgbotapi.Update) {
  orderId, err := strconv.Atoi(update.Message.CommandArguments())
  check(err)
  order, err := dao.GetOrder(string(update.Message.Chat.ID), orderId)
  var response string
  responseData := templateData {
    "orderName": order.Name,
    "orderId": order.Id,
    //"orderItems": order
  }
  if err == nil {
    response = getResponseText("order_result", responseData)
  } else {
    response = getResponseText("order_result_fail", responseData)
  }
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}

func closeOrder(update *tgbotapi.Update) {
  orderId, err := strconv.Atoi(update.Message.CommandArguments())
  check(err)
  err = dao.DeleteOrder(string(update.Message.Chat.ID), string(update.Message.From.ID), orderId)
  var response string
  responseData := templateData {
    "orderId": orderId,
  }
  if err == nil {
    response = getResponseText("close_success", responseData)
  } else {
    response = getResponseText("close_fail_order_id", responseData)
  }
  msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
  msg.ParseMode = "Markdown"
  telegramBot.Send(msg)
}
