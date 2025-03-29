package models

type PoliceDepartment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	WorkDays    string `json:"work_days"`
	WorkTime    string `json:"work_time"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	CreatedAt   string `json:"created_at"`
}
