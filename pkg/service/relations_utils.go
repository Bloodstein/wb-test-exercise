package service

func (r *RelationsService) OfficeAlreadyExists(officeId int) bool {

	return r.repo.RelationsRepository.OfficeAlreadyExists(officeId)
}

func (r *RelationsService) TelegramChatAlreadyExists(telegramChatId int) bool {

	return r.repo.RelationsRepository.TelegramChatAlreadyExists(telegramChatId)
}
