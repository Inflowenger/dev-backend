package models

type ContextRecord struct {
	ID        string         `json:"id"`
	Title     string         `json:"title"`
	Context   map[string]any `json:"context"`
	Header    map[string]any `json:"header"`
	UpdatedAt int64          `json:"updatedAt"`
	CreatedAt int64          `json:"createdAt"`
}

type ContextDoc struct {
	Data   string         `json:"data"`
	Header map[string]any `json:"header"`
}
