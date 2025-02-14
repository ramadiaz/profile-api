package models

import "time"

type TreeURLs struct {
	ID   uint   `gorm:"primaryKey"`
	UUID string `gorm:"not null;unique;index"`

	Name        string `gorm:"not null"`
	ShortURL    string `gorm:"not null;unique;index"`
	OriginalURL string `gorm:"not null;index"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
