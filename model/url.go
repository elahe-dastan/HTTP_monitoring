package model

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	ID     int `gorm:"primaryKey;AUTO_INCREMENT"`
	UserID int
	URL    string `gorm:"not null"`
	Period int    `gorm:"not null"`
	Statuses []Status	`gorm:"foreignkey:URLID"`
}
