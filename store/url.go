package store

import (
	"log"

	"github.com/elahe-dastan/HTTP_monitoring/model"
	"gorm.io/gorm"
)

type SQLURL struct {
	DB *gorm.DB
}

func NewURL(d *gorm.DB) SQLURL {
	return SQLURL{DB: d}
}

// Creates a table in the database that matches the URL table.
func (u SQLURL) Create() {
	if err := u.DB.Migrator().DropTable(&model.URL{}); err != nil {
		log.Fatal(err)
	}

	if err := u.DB.Migrator().CreateTable(&model.URL{}); err != nil {
		log.Fatal(err)
	}
}

func (u SQLURL) Insert(url model.URL) error {
	result := u.DB.Create(&url)

	return result.Error
}

func (u SQLURL) GetTable() ([]model.URL, error) {
	var models []model.URL

	result := u.DB.Find(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	return models, nil
}
