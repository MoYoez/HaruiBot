package HaruiBot

import (
	"log"
)

const (
	PanicLevel int = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

var (
	ConnectionLogger   = log.Logger{} // connection, receiving messages.
	BotSelfLogger      = log.Logger{} // bot information, Handler information.
	BotSelfPanicLogger = log.Logger{} // bot handler panic information yet.
)

func init() {
	// make init config sets here.
	ConnectionLogger.SetPrefix(Colorize("[Telegram-BotInstance]", FgBlue))
	BotSelfLogger.SetPrefix(Colorize("[HaruiBot-Events]", FgBlue))
	BotSelfPanicLogger.SetPrefix(Colorize("[HaruiBot-Panic]", FgRed)) // this will manually happen
	// when some workers not work yet in some instance.
}
