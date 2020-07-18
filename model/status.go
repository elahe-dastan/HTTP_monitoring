package model

import (
	"time"

	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	ID int `gorm:"primaryKey,AUTO_INCREMENT"`
	URLID      int `redis:"url"`
	Clock      time.Time	`redis:"clock"`
	StatusCode int	`redis:"status"`
}
