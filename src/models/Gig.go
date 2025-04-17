package models

import (
	"time"
)

type Gig struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Location    string    `gorm:"size:255;not null" json:"location"`
	Date        time.Time `json:"date"`
	Instrument  string    `gorm:"size:100;not null" json:"instrument"`
	Status      string    `gorm:"size:50;default:'open'" json:"status"`
}

type GigApplication struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	GigID      uint   `gorm:"not null" json:"gig_id"`
	Gig        Gig    `gorm:"foreignKey:GigID" json:"gig"`
	MusicianID uint   `gorm:"not null" json:"musician_id"`
	Musician   User   `gorm:"foreignKey:MusicianID" json:"-"`
	Status     string `gorm:"size:50;default:'pending'" json:"status"`
}
