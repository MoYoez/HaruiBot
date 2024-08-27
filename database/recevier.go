package database

import (
	"fmt"
	"github.com/MoYoez/HaruiBot/handler"
	"strconv"
)

func TransformUserConfig(updater handler.Updater) {
	// get user own information here, use updater.
	DB.Create(&UserModel{UserID: updater.SendFrom.ID}).Updates(&UserModel{
		UserID:      updater.SendFrom.ID,
		Username:    updater.SendFrom.UserName,
		FullName:    fmt.Sprintf("%s %s", updater.SendFrom.LastName, updater.SendFrom.FirstName),
		FromChannel: []string{strconv.FormatInt(updater.FromChat.ID, 10)}, // get this user chat id
		UserPhoto:   0,                                                    // get user photo.
	})
}
