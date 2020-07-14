package memory

import "github.com/gomodule/redigo/redis"

type Status struct {
	Redis redis.Conn
}

func NewStatus(r redis.Conn) Status {
	return Status{Redis:r}
}

