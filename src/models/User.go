package models

type User struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"size:255" json:"name"`
	Email      string `gorm:"unique;not null" json:"email"`
	Password   string `json:"-"`
	CognitoSub string `gorm:"uniqueIndex" json:"cognito_sub"`
	Instrument string `gorm:"size:100" json:"instrument"`
	Location   string `gorm:"size:255" json:"location"`
	Bio        string `gorm:"type:text" json:"bio"`
}
