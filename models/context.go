package models

type ContextRecord struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Context map[string]any `json:"context"`
	CreatedAt int64 `json:"createdAt"`
}