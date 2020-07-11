package store

import (
	"HTTP_monitoring/model"
	"database/sql"
	"log"
)

type SQLURL struct {
	DB      *sql.DB
}

func NewURL(d *sql.DB) SQLURL {
	return SQLURL{DB: d,
	}
}

// Creates a table in the database that matches the URL table and puts a trigger on it which deletes the
// rows that have expired after each insert
func (u SQLURL) Create() {
	_, err := u.DB.Exec("CREATE TABLE IF NOT EXISTS url (" +
		"id serial PRIMARY KEY," +
		"u INTEGER," +
		"url VARCHAR NOT NULL," +
		"FOREIGN KEY (u) REFERENCES users (id)" +
		");")
	if err != nil {
		log.Println("Cannot create url table due to the following error", err.Error())
	}
}

func (u SQLURL) Insert(url model.URL) error {
	_, err := u.DB.Exec("INSERT INTO url (u, url) VALUES ($1, $2)",
		url.UserId, url.Url)

	return err
}