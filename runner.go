package HaruiBot

import (
	"fmt"
	"github.com/MoYoez/HaruiBot/handler"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// GetInstance make a new instance bot here
func (register *BotRegister) GetInstance() *BotInstance {
	var NewBotInstance *tgbotapi.BotAPI
	for {
		OnReadyBotInstance, err := tgbotapi.NewBotAPIWithClient(register.Token, tgbotapi.APIEndpoint, &register.HttpClient)
		if err != nil {
			ConnectionLogger.Print(fmt.Print("Failed to connect Telegram Service, retry in 3 seconds."))
			time.Sleep(time.Second * 3)
			continue
		}
		NewBotInstance = OnReadyBotInstance
		break
	}

	if register.Debug {
		NewBotInstance.Debug = true // this will make information output as debugging msg here.
	}
	if register.BotName == "" {
		register.BotName = NewBotInstance.Self.UserName // bot only has username, no firstname / lastname ^^.
	}
	// if anything work, then print log for auth bot here.
	return &BotInstance{
		UnionId:           NewBotInstance.Self.ID,
		BotRegisterUpdate: register.OnUpdate,
		BotUserConfig:     register.BotConfig,
		Worker:            *NewBotInstance,
		Self:              NewBotInstance.Self,
	}
}

func (b *BotInstance) RunInstance() {
	u := tgbotapi.NewUpdate(0)
	u.AllowedUpdates = b.BotRegisterUpdate
	u.Timeout = 60 // set 60s timeout time.
	UpdaterChan := b.Worker.GetUpdatesChan(u)
	for getUpdater := range UpdaterChan {
		handler.HandleSense(getUpdater)
	}
}
