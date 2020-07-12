package db

import (
	"HTTP_monitoring/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq" //adding dialect for postgres
)

const DB = "postgres"

func New(config config.Database) *sql.DB {
	db, err := sql.Open(DB, config.Cstring())
	if err != nil {
		log.Printf("can not open connection to database due to the following err\n: %s", err)
	}

	return db
}
