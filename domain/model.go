package domain

type ModifyRequest struct {
	OfficeID       int `json:"officeId"`
	TelegramChatID int `json:"telegramChatId"`
}

type TelegramToOfficeRelation struct {
	ID             string `json:"id" bson:"_id"`
	OfficeID       int    `json:"officeId"`
	TelegramChatID int    `json:"telegramChatId"`
}
