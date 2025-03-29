package models

type EducationContent struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	MediaURL  string `json:"media_url"`
	CreatedAt string `json:"created_at"`
}
