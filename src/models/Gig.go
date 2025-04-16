package models

import (
	"time"
)

type Gig struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"` // Owner of the gig (Act/Band)
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Location    string    `gorm:"size:255;not null" json:"location"`
	Date        time.Time `json:"date"`
	Instrument  string    `gorm:"size:100;not null" json:"instrument"`  // Required instrument
	Status      string    `gorm:"size:50;default:'open'" json:"status"` // open, filled, closed
}

type GigApplication struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	GigID      uint   `gorm:"not null" json:"gig_id"`
	MusicianID uint   `gorm:"not null" json:"musician_id"`
	Status     string `gorm:"size:50;default:'pending'" json:"status"` // pending, accepted, rejected
}
