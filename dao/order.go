package dao

import (
  "public-order-bot/db"
  "public-order-bot/db/models"
)

func CreateOrder(userId string) {
  err := db.Conn.Insert(&models.Order{
    UserId: userId,
  })
  if err != nil {
    panic(err)
  }
}

func GetOrder(userId string) *models.Order {
  order := models.Order{UserId: userId}
  err := db.Conn.Select(&order)
  if err != nil {
    panic(err)
  }
  return &order
}
