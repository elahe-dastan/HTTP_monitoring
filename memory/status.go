package memory

import (
	"HTTP_monitoring/model"
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

type Status struct {
	Redis   redis.Conn
	Counter int
}

func NewStatus(r redis.Conn) Status {
	return Status{Redis: r,
		Counter: 0}
}

func (s *Status) Insert(status model.Status) {
	_, err := s.Redis.Do("HMSET", "status:" + strconv.Itoa(s.Counter), "url", status.URL, "clock",
		status.Clock, "status", status.StatusCode)
	if err != nil {
		log.Fatal(err)
	}

	s.Counter++
}

func (s *Status) Flush() []model.Status{
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

	return models
}