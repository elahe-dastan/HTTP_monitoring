package model

type URL struct {
	ID       int `gorm:"primaryKey;AUTO_INCREMENT"`
	UserID   int
	URL      string   `gorm:"not null"`
	Period   int      `gorm:"default:1"`
	Statuses []Status `gorm:"foreignkey:URLID"`
}
