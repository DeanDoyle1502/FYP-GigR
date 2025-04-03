package models

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:255"`
	Email      string `gorm:"unique;not null"`
	Password   string
	CognitoSub string `gorm:"uniqueIndex"`
	Instrument string `gorm:"size:100"`
	Location   string `gorm:"size:255"`
	Bio        string `gorm:"type:text"`
}
