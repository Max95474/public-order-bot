package dao

import (
  "public-order-bot/db"
	"public-order-bot/db/models"
)

func AddUser(userId string) (*models.User, error) {
  user := models.User{Id: userId}
  err := db.Conn.Select(&user)
  if err != nil {
    err = db.Conn.Insert(&user)
  }
	return &user, err
}

func GetUser(userId string) (*models.User, error) {
	user := models.User{Id: userId}
	err := db.Conn.Select(&user)
	return &user, err
}
