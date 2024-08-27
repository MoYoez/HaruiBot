package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Updater struct {
	UpdatedId   int64
	EventAction string
	Event       any
	FromChat    *tgbotapi.Chat
	SendFrom    *tgbotapi.User
}
