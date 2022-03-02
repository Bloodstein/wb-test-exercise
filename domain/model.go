package domain

type TelegramToOfficeRelation struct {
	OfficeID       string `json:"officeId"`
	TelegramChatID int    `json:"telegramChatId"`
}
