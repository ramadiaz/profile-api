package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type TreeURLOutput struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}
