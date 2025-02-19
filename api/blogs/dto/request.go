package dto

type Blogs struct {
	Title     string     `json:"title" validate:"required"`
	Content   string     `json:"content" validate:"required"`
	Tags      []BlogTags `json:"tags" validate:"required"`
	Thumbnail string     `json:"thumbnail" validate:"required,url"`
}

type BlogTags struct {
	Tag string `json:"tag" validate:"required"`
}

type FeaturedType string

const (
	Hot      FeaturedType = "hot"
	Featured FeaturedType = "featured"
)

type FeaturedBlogs struct {
	BlogUUID string       `json:"blog_uuid"`
	Type     FeaturedType `json:"type"`
}
