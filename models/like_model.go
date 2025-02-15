package models

import (
	"time"

	"gorm.io/gorm"
)

type Likes struct {
	gorm.Model

	ID        uint `gorm:"primaryKey"`

	IP        string
	Browser   string
	OS        string
	Device    string

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
