package dto

import "time"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type BlogOutput struct {
	UUID      string     `json:"uuid"`
	Slug      string     `json:"slug"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	Tags      []BlogTagOutput `json:"tags"`
}

type BlogTagOutput struct {
	Tag string `json:"tag" validate:"required"`
}
