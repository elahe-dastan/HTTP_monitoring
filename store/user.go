package store

import (
	"HTTP_monitoring/model"
	"database/sql"
	"errors"
	"log"
)


var ErrNotFound = errors.New("this user doesn't exist in the database")
var ErrWrongPass = errors.New("password is not correct")

type SQLUser struct {
	DB      *sql.DB
}

func NewUser(d *sql.DB) SQLUser {
	return SQLUser{DB: d,
	}
}

// Creates a table in the database that matches the User table.
func (u SQLUser) Create() {
	_, err := u.DB.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"id serial PRIMARY KEY," +
		"email VARCHAR NOT NULL," +
		"pass VARCHAR NOT NULL," +
		"CONSTRAINT email_unique UNIQUE (email)" +
		");")
	if err != nil {
		log.Println("Cannot create users table due to the following error", err.Error())
	}
}

func (u SQLUser) Insert(user model.User) error {
	_, err := u.DB.Exec("INSERT INTO users (email, pass) VALUES ($1, $2)",
		user.Email, user.Password)

	return err
}

func (u SQLUser) Retrieve(user model.User) (model.User, error) {
	var us model.User

	err := u.DB.QueryRow("SELECT * from users WHERE email = $1;", user.Email).Scan(
		&us.ID, &us.Email, &us.Password)
	if err != nil {
		log.Println(err)
	}

	if us.Email == "" {
		err = ErrNotFound
		return us, err
	}

	if us.Password != user.Password {
		err = ErrWrongPass
		return us, err
	}

	return us, err
}