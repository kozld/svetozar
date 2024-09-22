package models

type (
	Message struct {
		ID     int64  `gorm:"primaryKey" json:"id"`
		ChatID int64  `gorm:"primaryKey" json:"chatId"`
		Text   string `json:"text"`
		Status Status `json:"status"`
	}

	Status int8
)

const (
	MessageNew Status = iota
	MessageValid
	MessageInvalid
)
