package store

import (
	"errors"
	"log"

	"github.com/elahe-dastan/HTTP_monitoring/model"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("this user doesn't exist in the database")
var ErrWrongPass = errors.New("password is not correct")

type User interface {
	Insert(user model.User) error
	Retrieve(user model.User) (model.User, error)
}

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
}

func (u SQLUser) Insert(user model.User) error {
	result := u.DB.Create(&user)

	return result.Error
}

func (u SQLUser) Retrieve(user model.User) (model.User, error) {
	var us model.User

	u.DB.Where("email = ?", user.Email).First(&us)

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
