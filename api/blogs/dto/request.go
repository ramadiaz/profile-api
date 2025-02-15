package dto

type Blogs struct {
	Title   string     `json:"title" validate:"required"`
	Content string     `json:"content" validate:"required"`
	Tags    []BlogTags `json:"tags" validate:"required"`
}

type BlogTags struct {
	Tag string `json:"tag" validate:"required"`
}
