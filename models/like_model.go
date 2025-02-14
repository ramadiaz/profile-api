package models

import "time"

type Likes struct {
	ID        uint `gorm:"primaryKey"`

	IP        string
	Browser   string
	OS        string
	Device    string

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
