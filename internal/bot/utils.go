package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/internal/app"
)

// The function `getUserLanguage` returns the language code of the user from a Telegram update,
// defaulting to "en" if not available.
func GetUserLanguage(update tgbotapi.Update) string {
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
func ProcessCheckSubscription(app *app.App, update tgbotapi.Update) {
	lang := GetUserLanguage(update)

	var chatID int64

	var userID int64

	switch {
	case update.Message != nil:
		chatID = update.Message.Chat.ID
		userID = update.Message.From.ID

		msg := tgbotapi.NewMessage(chatID, app.LocaleStorage.GetText(lang, "checking_subscription"))
		SendMessage(app, msg)
	case update.CallbackQuery != nil:
		chatID = update.CallbackQuery.Message.Chat.ID
		userID = update.CallbackQuery.From.ID

		callbackResp := tgbotapi.NewCallback(
			update.CallbackQuery.ID,
			app.LocaleStorage.GetText(lang, "checking_subscription"),
		)

		_, err := app.Bot.Request(callbackResp)
		if err != nil {
			app.Logger.Error("❌ Error sending callback: %v", err)
		}
	default:
		app.Logger.Error("❌ Error: update contains neither Message nor CallbackQuery")

		return
	}

	isSubscribed, err := app.Subscription.IsUserSubscribed(userID)
	if err != nil {
		app.Logger.Error("❌ Error checking subscription: %v", err)

		msg := tgbotapi.NewMessage(chatID, app.LocaleStorage.GetText(lang, "error_message"))
		SendMessage(app, msg)

		return
	}

	var msg tgbotapi.MessageConfig
	if isSubscribed {
		msg = tgbotapi.NewMessage(chatID, app.LocaleStorage.GetText(lang, "subscribed_message"))
		msg.ReplyMarkup = GenerateSuccessKeyboard(app, lang)
	} else {
		msg = tgbotapi.NewMessage(chatID, app.LocaleStorage.GetText(lang, "not_subscribed_message"))
		msg.ReplyMarkup = GenerateKeyboard(app, lang)
	}

	SendMessage(app, msg)
}

// The function `processSuccess` sends a success message with an animation to a Telegram chat based on
// the type of update received.
func ProcessSuccess(app *app.App, update tgbotapi.Update) {
	lang := GetUserLanguage(update)

	var chatID int64

	switch {
	case update.CallbackQuery != nil:
		chatID = update.CallbackQuery.Message.Chat.ID

		callbackResp := tgbotapi.NewCallback(update.CallbackQuery.ID, app.LocaleStorage.GetText(lang, "access_is_open"))

		_, err := app.Bot.Request(callbackResp)
		if err != nil {
			app.Logger.Error("❌ Error sending callback: %v", err)
		}
	case update.Message != nil:
		chatID = update.Message.Chat.ID
	default:
		app.Logger.Error("❌ Error: update contains neither Message nor CallbackQuery")

		return
	}

	animation := tgbotapi.NewAnimation(chatID, tgbotapi.FileURL("https://i.imgur.com/MptNzLY.mp4"))

	_, err := app.Bot.Send(animation)
	if err != nil {
		app.Logger.Error("❌ Error sending animation (GIF mode): %v", err)
	}
}
