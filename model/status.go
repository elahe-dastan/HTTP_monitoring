package model

import (
	"time"
)

type Status struct {
	ID int `gorm:"primaryKey,AUTO_INCREMENT"`
	URLID      int `redis:"url"`
	Clock      time.Time	`redis:"clock"`
	StatusCode int	`redis:"status"`
}
