package models

type Message struct {
	ID         int    `json:"id"`
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Text       string `json:"text"`
	CreatedAt  string `json:"created_at"`
}
