package utils

import (
	"github.com/MoYoez/HaruiBot"
	"github.com/bytedance/sonic"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetChat(b *HaruiBot.BotInstance, Config *tgbotapi.ChatInfoConfig) (tgbotapi.ChatFullInfo, error) {
	getCaller, err := b.Worker.Request(Config)
	if err != nil {
		return tgbotapi.ChatFullInfo{}, err
	}
	var chat tgbotapi.ChatFullInfo
	err = sonic.Unmarshal(getCaller.Result, &chat)
	return chat, err
}

func GetFiles(b *HaruiBot.BotInstance, FileConfigID string) (str string, err error) {
	return b.Worker.GetFileDirectURL(FileConfigID)
}
