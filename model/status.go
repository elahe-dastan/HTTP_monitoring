package model

import (
	"time"
)

type Status struct {
	ID         int `gorm:"primaryKey,AUTO_INCREMENT"`
	URLID      int
	Clock      time.Time
	StatusCode int
}
