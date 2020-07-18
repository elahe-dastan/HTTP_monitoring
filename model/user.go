package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Urls     []URL	`gorm:"foreignkey:UserID"`
}
