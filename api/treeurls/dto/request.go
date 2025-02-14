package dto

type TreeURLs struct {
	Name        string `json:"name" validate:"required"`
	ShortURL    string `json:"short_url" validate:"required"`
	OriginalURL string `json:"original_url" validate:"required,url"`
}
