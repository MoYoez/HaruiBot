package utils

import (
	"github.com/MoYoez/HaruiBot"
	"github.com/bytedance/sonic"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *HaruiBot.BotInstance) Caller(Config tgbotapi.ChatInfoConfig) (tgbotapi.Chat, error) {
	getCaller, err := b.Worker.Request(Config)
	if err != nil {
		return tgbotapi.Chat{}, err
	}
	var chat tgbotapi.Chat
	err = sonic.Unmarshal(getCaller.Result, &chat)
	return chat, err
}
