package database

import (
	"fmt"
	"github.com/MoYoez/HaruiBot"
	"github.com/MoYoez/HaruiBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func SaveInformation(b *HaruiBot.BotInstance, updater tgbotapi.Update) {
	// get user own information here, use updater.
	DB.Create(&UserModel{UserID: updater.SentFrom().ID}).Updates(&UserModel{
		UserID:      updater.SentFrom().ID,
		Username:    updater.SentFrom().UserName,
		FullName:    fmt.Sprintf("%s %s", updater.SentFrom().LastName, updater.SentFrom().FirstName),
		FromChannel: []string{strconv.FormatInt(updater.FromChat().ID, 10)},                                                              // get this user chat id
		UserPhoto:   utils.GetAvatarImageID(b, &tgbotapi.ChatInfoConfig{ChatConfig: tgbotapi.ChatConfig{ChatID: updater.SentFrom().ID}}), // get a user photo (file id)
	})
}
