package db

import (
	"github.com/go-pg/pg"
	"public-order-bot/config"
)

var Conn *pg.DB

func init() {
	Conn = connect()
}

func connect() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     config.Config.Database.Username,
		Password: config.Config.Database.Password,
		Database: config.Config.Database.DatabaseName,
	})
}
