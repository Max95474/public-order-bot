package test

import (
  "testing"
  "public-order-bot/dao"
)

var orderId int

func TestCreateOrder(t *testing.T) {
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
