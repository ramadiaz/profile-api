package models

import (
	"time"

	"gorm.io/gorm"
)

type Blogs struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey"`
	UUID      string `gorm:"not null;unique;index"`
	Slug      string `gorm:"not null;unique;index"`
	Title     string `gorm:"not null"`
	Content   string `gorm:"not null"`
	Thumbnail string `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	Tags []BlogTags `gorm:"many2many:blog_tags_relation"`
}

type BlogTags struct {
	gorm.Model

	ID  uint   `gorm:"primaryKey"`
	Tag string `gorm:"not null;unique;index"`
}

type FeaturedType string

const (
	Hot      FeaturedType = "hot"
	Featured FeaturedType = "featured"
)

type FeaturedBlogs struct {
	gorm.Model

	ID       uint         `gorm:"primaryKey"`
	BlogUUID string       `gorm:"not null;index"`
	Type     FeaturedType `gorm:"type:featured_type;not null;index"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	Blog Blogs `gorm:"foreignKey:BlogUUID;references:UUID"`
}
