package store

import (
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
func (m SQLURL) Create() {
	_, err := m.DB.Exec("CREATE TABLE IF NOT EXISTS url (" +
		"id serial PRIMARY KEY," +
		"user INTEGER," +
		"url VARCHAR NOT NULL," +
		"FOREIGN KEY (user) REFERENCES user (id)" +
		");")
	if err != nil {
		log.Println("Cannot create map table due to the following error", err.Error())
	}
}