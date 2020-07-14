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

func (s Status) Insert(status model.Status) {
	_, err := s.Redis.Do("HMSET", "status:" + strconv.Itoa(s.Counter), "url", status.URL, "clock",
		status.Clock, "status", status.StatusCode)
	if err != nil {
		log.Fatal(err)
	}

	s.Counter++
}
