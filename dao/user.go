package dao

import (
  "public-order-bot/db"
	"public-order-bot/db/models"
)

func CreateUser(id string) {
	err := db.Conn.Insert(&models.User{
		Id: id,
	})
	if err != nil {
		panic(err)
	}
}

func GetUser(id string) *models.User {
	user := models.User{Id: id}
	err := db.Conn.Select(&user)
	if err != nil {
		panic(err)
	}
	return &user
}
