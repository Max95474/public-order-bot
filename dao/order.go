package dao

import (
  "public-order-bot/db"
  "public-order-bot/db/models"
)

func CreateOrder(chatId string, userId string, name string) (*models.Order, error) {
  _, err := AddUser(userId)
  newOrder := models.Order{ChatId: chatId, UserId: userId, Name: name}
  err = db.Conn.Insert(&newOrder)
  return &newOrder, err
}

func GetOrder(chatId string, orderId int) (*models.Order, error) {
  order := models.Order{ChatId: chatId, Id: orderId}
  err := db.Conn.Select(&order)
  return &order, err
}

func ListOrders(chatId string) ([]*models.Order, error) {
  var orders []*models.Order
  err := db.Conn.Model(&orders).
    Where("chat_id = ?", chatId).Select()
  return orders, err
}

func DeleteOrder(chatId string, userId string, orderId int) error {
  order := models.Order{ChatId: chatId, UserId: userId, Id: orderId}
  err := db.Conn.Delete(&order)
  return err
}
