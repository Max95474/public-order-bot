package main

import (
	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		_, err := db.Exec(`
				CREATE TABLE users (
					id VARCHAR(255) NOT NULL PRIMARY KEY
				);
				CREATE TABLE orders (
					id SERIAL PRIMARY KEY,
					user_id VARCHAR(255) NOT NULL REFERENCES users ON DELETE CASCADE
				);
				CREATE TABLE items (
					id SERIAL PRIMARY KEY,
					order_id INT NOT NULL REFERENCES orders ON DELETE CASCADE,
					text VARCHAR(255) NOT NULL
				);`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
				DROP TABLE items;
				DROP TABLE orders;
				DROP TABLE users`)
		return err
	})
}
