package dao

import (
  "public-order-bot/db"
  "public-order-bot/db/models"
  "errors"
)

func CreateItem(chatId string, orderId int, text string) (*models.Item, error) {
  order := models.Order{Id: orderId}
  err := db.Conn.Select(&order)
  if order.ChatId != chatId {
    return nil, errors.New("this order is not from your chat")
  }

  item := models.Item{
    OrderId: orderId,
    Text: text,
  }
  err = db.Conn.Insert(&item)
  return &item, err
}

func GetItems(chatId string, orderId int) ([]models.Item, error) {
  order := models.Order{Id: orderId}
  err := db.Conn.Select(&order)
  if order.ChatId != chatId {
    return nil, errors.New("this order is not from your chat")
  }

  var items []models.Item
  err = db.Conn.Model(&items).
      Where("order_id = ?", orderId).Select()

  if err != nil {
    panic(err)
  }
  return items, err
}
