package status

import (
	"log"
	"strconv"
	"time"

	"github.com/elahe-dastan/HTTP_monitoring/model"
	"github.com/gomodule/redigo/redis"
)

type RedisStatus struct {
	Redis   redis.Conn
	Counter int
}

func NewRedisStatus(r redis.Conn) RedisStatus {
	return RedisStatus{Redis: r,
		Counter: 0}
}

func (s *RedisStatus) Insert(status model.Status) {
	_, err := s.Redis.Do("HMSET", "status:"+strconv.Itoa(s.Counter), "url", status.URLID, "clock",
		status.Clock.Format(time.RFC3339), "status", status.StatusCode)
	if err != nil {
		log.Fatal(err)
	}

	s.Counter++
}

//nolint: gofumpt
func (s *RedisStatus) Flush() []model.Status {
	models := make([]model.Status, s.Counter)

	for i := 0; i < s.Counter; i++ {
		values, err := redis.Values(s.Redis.Do("HGETALL", "status:"+strconv.Itoa(i)))
		if err != nil {
			log.Fatal(err)
		}

		var status redisStatus

		if err := redis.ScanStruct(values, &status); err != nil {
			log.Fatal(err)
		}

		models[i].URLID = status.URLID
		models[i].Clock, err = time.Parse(time.RFC3339, status.Clock)

		if err != nil {
			log.Fatal(err)
		}

		models[i].StatusCode = status.StatusCode

		_, err = s.Redis.Do("DEL", "status:"+strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
	}

	s.Counter = 0

	return models
}
