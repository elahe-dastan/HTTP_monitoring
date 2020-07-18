package db

import (
	"log"

	"github.com/elahe-dastan/HTTP_monitoring/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//nolint: gofumpt
func New(config config.Database) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Cstring()), &gorm.Config{})
	if err != nil {
		log.Printf("can not open connection to database due to the following err\n: %s", err)
	}

	return db
}
