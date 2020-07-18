package model

type URL struct {
	ID     int `gorm:"primaryKey;AUTO_INCREMENT"`
	User   User
	URL    string `gorm:"not null"`
	Period int    `gorm:"not null"`
}
