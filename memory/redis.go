package memory

import (
	"HTTP_monitoring/config"
	"log"

	"github.com/gomodule/redigo/redis"
)

func New(config config.Redis) redis.Conn {
	conn, err := redis.Dial("tcp", config.Host + ":" + config.Port)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	return conn
}
