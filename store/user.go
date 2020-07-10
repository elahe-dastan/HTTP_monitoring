package store

import (
	"HTTP_monitoring/model"
	"database/sql"
	"log"
)

type SQLUser struct {
	DB      *sql.DB
}

func NewUser(d *sql.DB) SQLUser {
	return SQLUser{DB: d,
	}
}

// Creates a table in the database that matches the User table and puts a trigger on it which deletes the
// rows that have expired after each insert
func (u SQLUser) Create() {
	_, err := u.DB.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"id serial PRIMARY KEY," +
		"email VARCHAR NOT NULL," +
		"pass VARCHAR NOT NULL," +
		"CONSTRAINT email_unique UNIQUE (email)" +
		");")
	if err != nil {
		log.Println("Cannot create user table due to the following error", err.Error())
	}
}

func (u SQLUser) Insert(user model.User) error {
	_, err := u.DB.Exec("INSERT INTO users (email, pass) VALUES ($1, $2)",
		user.Email, user.Password)

	return err
}