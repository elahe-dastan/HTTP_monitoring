package store

import (
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
	u.DB.Migrator().CreateTable(&model.URL{})
	//_, err := u.DB.Exec("CREATE TABLE IF NOT EXISTS url (" +
	//	"id serial PRIMARY KEY," +
	//	"u INTEGER," +
	//	"url VARCHAR NOT NULL," +
	//	"period INTEGER," +
	//	"FOREIGN KEY (u) REFERENCES users (id)" +
	//	");")
	//if err != nil {
	//	log.Println("Cannot create url table due to the following error", err.Error())
	//}
}

func (u SQLURL) Insert(url model.URL) {
	u.DB.Create(&url)
	//_, err := u.DB.Exec("INSERT INTO url (u, url, period) VALUES ($1, $2, $3)",
	//	url.UserID, url.URL, url.Period)
	//
	//return err
}

func (u SQLURL) GetTable() {
	var user model.User

	u.DB.Find(&user)
	//rows, err := u.DB.Find()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//return rows
}
