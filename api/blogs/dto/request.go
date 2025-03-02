package dto

type Blogs struct {
	Title     string     `json:"title" validate:"required"`
	Content   string     `json:"content" validate:"required"`
	Tags      []BlogTags `json:"tags" validate:"required"`
	Thumbnail string     `json:"thumbnail" validate:"required,url"`
	Text      string     `json:"text" validate:"required"`
}

type BlogUpdate struct {
	UUID      string     `json:"uuid" validate:"required,uuid4"`
	Title     string     `json:"title" validate:"required"`
	Content   string     `json:"content" validate:"required"`
	Tags      []BlogTags `json:"tags" validate:"required"`
	Thumbnail string     `json:"thumbnail" validate:"required,url"`
	Text      string     `json:"text" validate:"required"`
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
	BlogUUID string       `json:"blog_uuid" validate:"required,uuid4"`
	Type     FeaturedType `json:"type" validate:"required,oneof=hot featured"`
}
