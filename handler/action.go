package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// HandleSense when receiving telegram update, handle and go on sence service.
func HandleSense(Handler tgbotapi.Update) (actionMain tgbotapi.Update, EventAction string, Raw any) {
	// On Update Message have some situations
	switch {
	case Handler.Message != nil:
		return Handler, "MessageEvent", Handler.Message // go match message action.
	case Handler.EditedMessage != nil:
		return Handler, "MessageModifiedEvent", Handler.EditedMessage // when handling with an edited message, some works should be updated when data refreshing.
	case Handler.CallbackQuery != nil:
		return Handler, "CallBackQueryEvent", Handler.CallbackQuery
	case Handler.InlineQuery != nil:
		return Handler, "InlineQueryEvent", Handler.InlineQuery
	case Handler.ChosenInlineResult != nil:
		return Handler, "ChosenInlineQueryEvent", Handler.ChosenInlineResult
	case Handler.ChannelPost != nil:
		return Handler, "ChannelPostEvent", Handler.ChannelPost
	case Handler.EditedChannelPost != nil:
		return Handler, "ChannelPostModifiedEvent", Handler.EditedChannelPost
	case Handler.ChatBoost != nil:
		return Handler, "ChannelBoostEvent", Handler.ChatBoost
	case Handler.ChatBoostRemoved != nil:
		return Handler, "ChannelBoostRemoveEvent", Handler.ChatBoostRemoved
	case Handler.ChatJoinRequest != nil:
		return Handler, "ChatMemberJoinRequestEvent", Handler.ChatJoinRequest // be aware that that's only chat request.
	case Handler.ChatMember != nil:
		return Handler, "ChatMemberEvent", Handler.ChatMember
	case Handler.MyChatMember != nil:
		return Handler, "MyChatMemberEvent", Handler.MyChatMember
	case Handler.MessageReaction != nil:
		return Handler, "MessageReactionEvent", Handler.MessageReaction
	case Handler.MessageReactionCount != nil:
		return Handler, "MessageReactionChannelEvent", Handler.MessageReactionCount // this will always happen in Channel self.
		// TODO: POLL and awesome.
	}
	return Handler, "UnknownEvent", nil
}

func HandlerModify(actionMain tgbotapi.Update, EventAction string, Raw any) Updater {
	return Updater{
		UpdatedId:   int64(actionMain.UpdateID),
		EventAction: EventAction,
		Event:       Raw,
		FromChat:    actionMain.FromChat(),
		SendFrom:    actionMain.SentFrom(),
	}
}
