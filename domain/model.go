package domain

type TelegramToOfficeRelation struct {
	ID             int    `json:"id"`
	OfficeID       string `json:"officeId"`
	TelegramChatID int    `json:"telegramChatId"`
}
