package main

import (
	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		_, err := db.Exec(`
				CREATE TABLE "user" (
					id VARCHAR(255) NOT NULL PRIMARY KEY
				);
				CREATE TABLE "order" (
					id SERIAL PRIMARY KEY,
					user_id VARCHAR(255) NOT NULL REFERENCES "user" ON DELETE CASCADE
				);
				CREATE TABLE item (
					id SERIAL PRIMARY KEY,
					order_id INT NOT NULL REFERENCES "order" ON DELETE CASCADE,
					text VARCHAR(255) NOT NULL
				);`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
				DROP TABLE item;
				DROP TABLE "order";
				DROP TABLE "user"`)
		return err
	})
}
