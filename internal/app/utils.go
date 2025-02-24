package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// The function `getUserLanguage` returns the language code of the user from a Telegram update,
// defaulting to "en" if not available.
func (a *App) GetUserLanguage(update tgbotapi.Update) string {
	if update.Message != nil {
		return update.Message.From.LanguageCode
	}

	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.LanguageCode
	}

	return "en"
}

// The function `processCheckSubscription` checks a user's subscription status and sends a
// corresponding message to the user.
func (a *App) ProcessCheckSubscription(update tgbotapi.Update) {
	lang := a.GetUserLanguage(update)

	var chatID int64

	var userID int64

	switch {
	case update.Message != nil:
		chatID = update.Message.Chat.ID
		userID = update.Message.From.ID

		msg := tgbotapi.NewMessage(chatID, a.LocaleStorage.GetText(lang, "checking_subscription"))
		a.SendMessage(msg)
	case update.CallbackQuery != nil:
		chatID = update.CallbackQuery.Message.Chat.ID
		userID = update.CallbackQuery.From.ID

		callbackResp := tgbotapi.NewCallback(
			update.CallbackQuery.ID,
			a.LocaleStorage.GetText(lang, "checking_subscription"),
		)

		_, err := a.Bot.Request(callbackResp)
		if err != nil {
			a.Logger.Error("❌ Error sending callback: %v", err)
		}
	default:
		a.Logger.Error("❌ Error: update contains neither Message nor CallbackQuery")

		return
	}

	isSubscribed, err := a.Subscription.IsUserSubscribed(userID)
	if err != nil {
		a.Logger.Error("❌ Error checking subscription: %v", err)

		msg := tgbotapi.NewMessage(chatID, a.LocaleStorage.GetText(lang, "error_message"))
		a.SendMessage(msg)

		return
	}

	var msg tgbotapi.MessageConfig
	if isSubscribed {
		msg = tgbotapi.NewMessage(chatID, a.LocaleStorage.GetText(lang, "subscribed_message"))
		msg.ReplyMarkup = a.GenerateSuccessKeyboard(lang)
	} else {
		msg = tgbotapi.NewMessage(chatID, a.LocaleStorage.GetText(lang, "not_subscribed_message"))
		msg.ReplyMarkup = a.GenerateKeyboard(lang)
	}

	a.SendMessage(msg)
}

// The function `processSuccess` sends a success message with an animation to a Telegram chat based on
// the type of update received.
func (a *App) ProcessSuccess(update tgbotapi.Update) {
	lang := a.GetUserLanguage(update)

	var chatID int64

	switch {
	case update.CallbackQuery != nil:
		chatID = update.CallbackQuery.Message.Chat.ID

		callbackResp := tgbotapi.NewCallback(update.CallbackQuery.ID, a.LocaleStorage.GetText(lang, "access_is_open"))

		_, err := a.Bot.Request(callbackResp)
		if err != nil {
			a.Logger.Error("❌ Error sending callback: %v", err)
		}
	case update.Message != nil:
		chatID = update.Message.Chat.ID
	default:
		a.Logger.Error("❌ Error: update contains neither Message nor CallbackQuery")

		return
	}

	animation := tgbotapi.NewAnimation(chatID, tgbotapi.FileURL("https://i.imgur.com/MptNzLY.mp4"))

	_, err := a.Bot.Send(animation)
	if err != nil {
		a.Logger.Error("❌ Error sending animation (GIF mode): %v", err)
	}
}
