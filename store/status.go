package store

import (
	"log"
	"strconv"

	"github.com/elahe-dastan/HTTP_monitoring/model"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

type SQLStatus struct {
	DB *gorm.DB
}

type RedisStatus struct {
	Redis   redis.Conn
	Counter int
}

func NewSQLStatus(d *gorm.DB) SQLStatus {
	return SQLStatus{DB: d}
}

// Creates a table in the database that matches the status table and puts a trigger on it which deletes the
// rows that have expired after each insert.
func (m SQLStatus) Create() {
	m.DB.Migrator().CreateTable(&model.Status{})
	//_, err := m.DB.Exec("CREATE TABLE IF NOT EXISTS status (" +
	//	"id serial PRIMARY KEY," +
	//	"url INTEGER," +
	//	"clock TIMESTAMP NOT NULL," +
	//	"status INTEGER NOT NULL," +
	//	"FOREIGN KEY (url) REFERENCES url (id)" +
	//	");")
	//if err != nil {
	//	log.Println("Cannot create status table due to the following error", err.Error())
	//}

	m.DB.Exec("create or replace function delete_expired_row() " +
		"returns trigger as " +
		"$BODY$ " +
		"begin " +
		"delete from status where clock < NOW() - INTERVAL '2 days'; " +
		"return null; " +
		"end; " +
		"$BODY$ " +
		"LANGUAGE plpgsql;" +
		"create trigger delete_expired_rows " +
		"after insert " +
		"on status " +
		"for each row " +
		"execute procedure delete_expired_row();")

	//_, err = m.DB.Exec()
	//
	//if err != nil {
	//	log.Println("Cannot create put trigger on status table due to the following error", err.Error())
	//}
}

func (m SQLStatus) Insert(status model.Status) {
	//date, err := time.Parse(time.Stamp, status.Clock)
	//if err != nil {
	//	return err
	//}
	m.DB.Create(&status)
	//_, err = m.DB.Exec("INSERT INTO status (url, clock, status) VALUES ($1, $2, $3)",
	//	status.URL, date, status.StatusCode)
	//
	//return err
}


func NewRedisStatus(r redis.Conn) RedisStatus {
	return RedisStatus{Redis: r,
		Counter: 0}
}

func (s *RedisStatus) Insert(status model.Status) {
	_, err := s.Redis.Do("HMSET", "status:" + strconv.Itoa(s.Counter), "url", status.URL, "clock",
		status.Clock, "status", status.StatusCode)
	if err != nil {
		log.Fatal(err)
	}

	s.Counter++
}

func (s *RedisStatus) Flush() []model.Status{
	models := make([]model.Status, s.Counter)

	for i := 0; i < s.Counter; i++ {
		values, err := redis.Values(s.Redis.Do("HGETALL", "status:" + strconv.Itoa(i)))
		if err != nil {
			log.Fatal(err)
		}

		var status model.Status
		err = redis.ScanStruct(values, &status)
		if err != nil {
			log.Fatal(err)
		}

		models[i] = status

		_, err = s.Redis.Do("DEL", "status:" + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
	}

	s.Counter = 0

	return models
}
