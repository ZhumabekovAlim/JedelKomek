package models

type EmergencyCall struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	CreatedAt string `json:"created_at"`
}
