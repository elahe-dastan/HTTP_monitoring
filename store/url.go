package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/elahe-dastan/HTTP_monitoring/model"
)

type SQLURL struct {
	DB *sql.DB
}

func NewURL(d *sql.DB) SQLURL {
	return SQLURL{DB: d}
}

// Creates a table in the database that matches the URL table.
func (u SQLURL) Create() {
	_, err := u.DB.Exec("CREATE TABLE IF NOT EXISTS url (" +
		"id serial PRIMARY KEY," +
		"u INTEGER," +
		"url VARCHAR NOT NULL," +
		"period INTEGER," +
		"FOREIGN KEY (u) REFERENCES users (id)" +
		");")
	if err != nil {
		log.Println("Cannot create url table due to the following error", err.Error())
	}
}

func (u SQLURL) Insert(url model.URL) error {
	_, err := u.DB.Exec("INSERT INTO url (u, url, period) VALUES ($1, $2, $3)",
		url.UserID, url.URL, url.Period)

	return err
}

func (u SQLURL) GetTable() *sql.Rows {
	rows, err := u.DB.Query("SELECT * FROM url")
	if err != nil {
		fmt.Println(err)
	}

	return rows
}
