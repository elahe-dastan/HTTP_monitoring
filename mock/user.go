package mock

import (
	"errors"

	"github.com/elahe-dastan/HTTP_monitoring/model"
)

//nolint: gofumpt
var ErrDuplicateEmail = errors.New("this email exists")
var ErrWrongPass = errors.New("password is not correct")
var ErrNotFound = errors.New("this user doesn't exist in the database")

type User struct {
	Info map[string]string
}

func (u User) Insert(user model.User) error {
	_, ok := u.Info[user.Email]
	if ok {
		return ErrDuplicateEmail
	}

	u.Info[user.Email] = user.Password

	return nil
}

func (u User) Retrieve(user model.User) (model.User, error) {
	pass, ok := u.Info[user.Email]

	if ok {
		if user.Password == pass {
			return model.User{
				ID:       0,
				Email:    user.Email,
				Password: pass,
				Urls:     nil,
			}, nil
		}

		return model.User{}, ErrWrongPass
	}

	return model.User{}, ErrNotFound
}
