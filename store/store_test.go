package store

import (
	"testing"

	"github.com/elahe-dastan/HTTP_monitoring/config"
	"github.com/elahe-dastan/HTTP_monitoring/db"
	"github.com/elahe-dastan/HTTP_monitoring/model"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	cfg := config.Read()
	d := db.New(cfg.Database)
	user := NewUser(d)

	assert.Nil(t, user.Insert(model.User{
		Email:    "parham.alvani@gmail.com",
		Password: "1373",
	}))

	user.Retrieve()
}