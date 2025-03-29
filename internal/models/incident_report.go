package models

type IncidentReport struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
	MediaURL    string `json:"media_url"`
	TypeID      string `json:"type_id"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	CreatedAt   string `json:"created_at"`
}
