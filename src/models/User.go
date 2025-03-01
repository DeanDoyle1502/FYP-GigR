package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
