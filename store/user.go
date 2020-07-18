package store

import (
	"errors"
	"log"

	"github.com/elahe-dastan/HTTP_monitoring/model"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("this user doesn't exist in the database")
var ErrWrongPass = errors.New("password is not correct")

type SQLUser struct {
	DB *gorm.DB
}

func NewUser(d *gorm.DB) SQLUser {
	return SQLUser{DB: d}
}

// Creates a table in the database that matches the User table.
func (u SQLUser) Create() {
	if err := u.DB.Migrator().DropTable(&model.User{}); err != nil {
		log.Fatal(err)
	}

	if err := u.DB.Migrator().CreateTable(&model.User{}); err != nil {
		log.Fatal(err)
	}
	//_, err := u.DB.Exec("CREATE TABLE IF NOT EXISTS users (" +
	//	"id serial PRIMARY KEY," +
	//	"email VARCHAR NOT NULL," +
	//	"pass VARCHAR NOT NULL," +
	//	"CONSTRAINT email_unique UNIQUE (email)" +
	//	");")
	//if err != nil {
	//	log.Println("Cannot create users table due to the following error", err.Error())
	//}
}

func (u SQLUser) Insert(user model.User) {
	u.DB.Create(&user)
	//check its primary key
	//_, err := u.DB.Exec("INSERT INTO users (email, pass) VALUES ($1, $2)",
	//	user.Email, user.Password)

	//return err
}

//nolint: gofumpt
func (u SQLUser) Retrieve(user model.User) (model.User, error) {
	var us model.User

	u.DB.Where("email = ?", user.Email).First(&us)
	//if err != nil {
	//	log.Println(err)
	//}

	var err error

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
