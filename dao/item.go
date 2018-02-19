package dao

import (
  "public-order-bot/db"
  "public-order-bot/db/models"
)

func CreateItem(orderId int, text string) {
  err := db.Conn.Insert(&models.Item{
    OrderId: orderId,
    Text: text,
  })
  if err != nil {
    panic(err)
  }
}

func GetItems(orderId int) []models.Item {
  var items []models.Item
  err := db.Conn.Model(&items).
      Where("order_id = ?", orderId).Select()

  if err != nil {
    panic(err)
  }
  return items
}
