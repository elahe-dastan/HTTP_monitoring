package store_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/elahe-dastan/HTTP_monitoring/config"
	"github.com/elahe-dastan/HTTP_monitoring/db"
	"github.com/elahe-dastan/HTTP_monitoring/memory"
	"github.com/elahe-dastan/HTTP_monitoring/model"
	"github.com/elahe-dastan/HTTP_monitoring/store"
	"github.com/elahe-dastan/HTTP_monitoring/store/status"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	cfg := config.Read()
	d := db.New(cfg.Database)
	user := store.NewUser(d)

	m := model.User{
		Email:    "parham.alvani@gmail.com",
		Password: "1373",
	}

	assert.Nil(t, user.Insert(m))

	u, err := user.Retrieve(m)

	assert.Nil(t, err)

	assert.Equal(t, m.Email, u.Email)
	assert.Equal(t, m.Password, u.Password)
}

func TestURL(t *testing.T) {
	cfg := config.Read()
	d := db.New(cfg.Database)
	user := store.NewUser(d)

	m := model.User{
		ID:       1,
		Email:    "elahe.dstn@gmail.com",
		Password: "1373",
	}

	if err := user.Insert(m); err != nil {
		fmt.Println(err)
	}

	url := store.NewURL(d)

	u := model.URL{
		UserID: 1,
		URL:    "https://www.google.com",
		Period: 2,
	}

	assert.Nil(t, url.Insert(u))

	_, err := url.GetTable()
	assert.Nil(t, err)
}

func TestRedis(t *testing.T) {
	cfg := config.Read()
	r := memory.New(cfg.Redis)
	redis := status.NewRedisStatus(r)

	m := model.Status{
		URLID:      1,
		Clock:      time.Now(),
		StatusCode: 200,
	}

	redis.Insert(m)

	statuses := redis.Flush()
	st := statuses[0]

	assert.Equal(t, m.URLID, st.URLID)
	assert.Equal(t, m.Clock.Format(time.RFC3339), st.Clock.Format(time.RFC3339))
	assert.Equal(t, m.StatusCode, st.StatusCode)
}
