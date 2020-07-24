package status

import (
	"log"

	"github.com/elahe-dastan/HTTP_monitoring/model"
	"gorm.io/gorm"
)

type Status interface {
	Insert(status model.Status) error
}

type SQLStatus struct {
	DB *gorm.DB
}

func NewSQLStatus(d *gorm.DB) SQLStatus {
	return SQLStatus{DB: d}
}

// Creates a table in the database that matches the status table and puts a trigger on it which deletes the
// rows that have expired after each insert.
func (m SQLStatus) Create() {
	if err := m.DB.Migrator().DropTable(&model.Status{}); err != nil {
		log.Fatal(err)
	}

	if err := m.DB.Migrator().CreateTable(&model.Status{}); err != nil {
		log.Fatal(err)
	}

	m.DB.Exec("create or replace function delete_expired_row() " +
		"returns trigger as " +
		"$BODY$ " +
		"begin " +
		"delete from statuses where clock < NOW() - INTERVAL '2 days'; " +
		"return null; " +
		"end; " +
		"$BODY$ " +
		"LANGUAGE plpgsql;" +
		"create trigger delete_expired_rows " +
		"after insert " +
		"on statuses " +
		"for each row " +
		"execute procedure delete_expired_row();")
}

func (m SQLStatus) Insert(status model.Status) error {
	result := m.DB.Create(&status)

	return result.Error
}
