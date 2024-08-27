package HaruiBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
)

type BotRegister struct {
	BotConfig
	Token      string
	Debug      bool
	HttpClient http.Client // make request work with proxy if you are in Mainland China or some countries that you cannot contact.
	OnUpdate   []string    // Telegram requests for updated information, which can ensure what bot will receive data here.
}

type BotConfig struct {
	BotName         string  // if BotName is null, them use Telegram Provided.
	AdministratorID []int64 // if null, them admin permission will not work yet.
}

// BotInstance bot data will store into mapped runner.
type BotInstance struct {
	UnionId           int64 // union id is generated when an instance work here, to ensure instance worker can be found by id.
	BotUserConfig     BotConfig
	BotRegisterUpdate []string
	Worker            tgbotapi.BotAPI
	Self              tgbotapi.User
}
