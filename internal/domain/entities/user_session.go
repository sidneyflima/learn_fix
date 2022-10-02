package entities

type UserSession struct {
	ID           string `json:"id,omitempty"`
	SenderCompID string `json:"senderCompID,omitempty"`
}
