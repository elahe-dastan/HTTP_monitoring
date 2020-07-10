package store

import (
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
	_, err := u.DB.Exec("CREATE TABLE IF NOT EXISTS user (" +
		"id serial PRIMARY KEY," +
		"email VARCHAR NOT NULL," +
		"pass VARCHAR NOT NULL" +
		");")
	if err != nil {
		log.Println("Cannot create map table due to the following error", err.Error())
	}
}