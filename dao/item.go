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

//func GetItems(orderId int) []*models.Item {
//  item := models.Item{OrderId: orderId}
//  err := db.Conn.Select(&item)
//  if err != nil {
//    panic(err)
//  }
//  return &item
//}
