package models

import "time"

type Incognitos struct {
	ID      uint   `gorm:"primaryKey"`
	UUID    string `gorm:"not null;unique;index"`
	Message string `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}