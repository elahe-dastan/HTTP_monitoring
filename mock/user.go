package mock

import (
	"errors"

	"github.com/elahe-dastan/HTTP_monitoring/model"
)

type User struct {
	info map[string]string
}

func (u *User) Insert(user model.User) error {
	_, ok := u.info[user.Email]
	if ok {
		return errors.New("this email exists")
	}

	u.info[user.Email] = user.Password

	return nil
}

func (u *User) Retrieve(user model.User) (model.User, error)  {
	pass, ok := u.info[user.Email]

	if ok {
		if user.Password == pass {
			return model.User{
				ID:       0,
				Email:    user.Email,
				Password: pass,
				Urls:     nil,
			}, nil
		}

		return model.User{}, errors.New("password is wrong")
	}

	return model.User{}, errors.New("email does not exist")
}
