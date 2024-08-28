package utils

import (
	"github.com/MoYoez/HaruiBot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"image"
	"net/http"
)

func GetAvatarImageID(b *HaruiBot.BotInstance, ChatConfig *tgbotapi.ChatInfoConfig) string {
	NewChat, err := GetChat(b, ChatConfig)
	if err != nil {
		return ""
	}
	return NewChat.Photo.SmallFileID
}

func GetAvatarImage(b *HaruiBot.BotInstance, ChatConfig *tgbotapi.ChatInfoConfig) (Result bool, img image.Image) {
	NewChat, err := GetChat(b, ChatConfig)
	if err != nil {
		return false, nil
	}
	chatPhoto := NewChat.Photo.SmallFileID
	chatPhotoDirectLink, err := GetFiles(b, chatPhoto)
	if err != nil {
		return false, nil
	}
	datas, err := http.Get(chatPhotoDirectLink)
	// avatar
	avatarByteUni, _, _ := image.Decode(datas.Body)
	return true, avatarByteUni
}
