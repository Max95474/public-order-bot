package test

import (
  "testing"
  "public-order-bot/dao"
  "public-order-bot/db/models"
)

var orderId int

func TestOrder(t *testing.T) {
  // Create user
  user, err := dao.AddUser("123123")
  if err != nil {
    t.Failed()
  }
  user, err = dao.GetUser(user.Id)
  if err != nil {
    t.Failed()
  }
  if user.Id != "123123" {
    t.Failed()
  }

  // Create order
  newOrder, err := dao.CreateOrder("5555", "123123", "test order")
  if err != nil {
    t.Error(err)
  }
  if newOrder.ChatId != "5555" ||
    newOrder.UserId != "123123" ||
      newOrder.Name != "test order" {
    t.Failed()
  }
  orderId = newOrder.Id

  // Add item
  var item *models.Item
  item, err = dao.CreateItem("5555", orderId, "Item text goes here")
  if item == nil {
    t.Failed()
  }
  var otherItem *models.Item
  otherItem, err = dao.CreateItem("7777", orderId, "Item text goes here")
  if otherItem != nil {
    t.Error(err)
  }
}

func TestGetOrder(t *testing.T) {
  order, err := dao.GetOrder("5555", orderId)
  if err != nil {
    t.Error(err)
  }
  if order.Name != "test order" {
    t.Failed()
  }
}

func TestGetOrderList(t *testing.T) {
  orders, err := dao.ListOrders("5555")
  if err != nil {
    t.Error(err)
  }
  if len(orders) == 0 {
    t.Failed()
  }
  if orders[0].Name != "test order" {
    t.Failed()
  }
}

func TestDeleteOrder(t *testing.T) {
  err := dao.DeleteOrder("5555", "123123", orderId)
  if err != nil {
    t.Error(err)
  }
}
