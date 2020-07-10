package store

import (
	"database/sql"
	"log"
)


type SQLStatus struct {
	DB      *sql.DB
}

func NewStatus(d *sql.DB) SQLStatus {
	return SQLStatus{DB: d,
	}
}

// Creates a table in the database that matches the status table and puts a trigger on it which deletes the
// rows that have expired after each insert
func (m SQLStatus) Create() {
	_, err := m.DB.Exec("CREATE TABLE IF NOT EXISTS status (" +
		"id serial PRIMARY KEY," +
		"url INTEGER," +
		"clock TIMESTAMP NOT NULL," +
		"status INTEGER NOT NULL," +
		"FOREIGN KEY (url) REFERENCES url (id)" +
		");")
	if err != nil {
		log.Println("Cannot create map table due to the following error", err.Error())
	}

	_, err = m.DB.Exec("create or replace function delete_expired_row() " +
		"returns trigger as " +
		"$BODY$ " +
		"begin " +
		"delete from status where clock < NOW() - INTERVAL '2 days'; " +
		"return null; " +
		"end; " +
		"$BODY$ " +
		"LANGUAGE plpgsql;" +
		"create trigger delete_expired_rows " +
		"after insert " +
		"on status " +
		"for each row " +
		"execute procedure delete_expired_row();")

	if err != nil {
		log.Println("Cannot create put trigger on map table due to the following error", err.Error())
	}
}