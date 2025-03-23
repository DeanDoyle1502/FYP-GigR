package models

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:255"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:""`
	CognitoSub string `gorm:"uniqueIndex"`
}
