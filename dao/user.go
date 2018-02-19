package dao

import (
  "public-order-bot/db"
	"public-order-bot/db/models"
)

func CreateUser(id string) (*models.User, error) {
  user := models.User{Id: id}
	err := db.Conn.Insert(&user)
	return &user, err
}

func GetUser(id string) (*models.User, error) {
	user := models.User{Id: id}
	err := db.Conn.Select(&user)
	return &user, err
}
