package db

import (
	"github.com/go-pg/pg"
	"errors"
)

var db *pg.DB

// Initialize database connection
func Init(database string, user string, password string) {
	if db == nil {
		db = pg.Connect(&pg.Options{
			User:     user,
			Password: password,
			Database: database,
		})
	}
}

// Get initialized database connection instance
func Get() (*pg.DB, error) {
	if db == nil {
		return nil, errors.New("initialize database first")
	}
	return db, nil
}
