package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/internal/locale"
	"github.com/vaaxooo/tgaccessbot/internal/services"
	"github.com/vaaxooo/tgaccessbot/pkg/logger"
)

// The function `getUserLanguage` returns the language code of the user from a Telegram update,
// defaulting to "en" if not available.
func getUserLanguage(update tgbotapi.Update) string {
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
func processCheckSubscription(bot *tgbotapi.BotAPI, update tgbotapi.Update, subService *services.SubscriptionService, channels []string, successRedirectURL string) {
	lang := getUserLanguage(update)

	var chatID int64
	var userID int64

	if update.Message != nil {
		chatID = update.Message.Chat.ID
		userID = update.Message.From.ID

		msg := tgbotapi.NewMessage(chatID, locale.GetText(lang, "checking_subscription"))
		sendMessage(bot, msg)
	} else if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.Message.Chat.ID
		userID = update.CallbackQuery.From.ID

		callbackResp := tgbotapi.NewCallback(update.CallbackQuery.ID, locale.GetText(lang, "checking_subscription"))
		_, err := bot.Request(callbackResp)
		if err != nil {
			logger.Error("❌ Error sending callback: %v", err)
		}
	} else {
		logger.Error("❌ Error: update contains neither Message nor CallbackQuery")
		return
	}

	isSubscribed, err := subService.IsUserSubscribed(int64(userID))
	if err != nil {
		logger.Error("❌ Error checking subscription: %v", err)
		msg := tgbotapi.NewMessage(chatID, locale.GetText(lang, "error_message"))
		sendMessage(bot, msg)
		return
	}

	var msg tgbotapi.MessageConfig
	if isSubscribed {
		msg = tgbotapi.NewMessage(chatID, locale.GetText(lang, "subscribed_message"))
		msg.ReplyMarkup = generateSuccessKeyboard(lang, successRedirectURL)
	} else {
		msg = tgbotapi.NewMessage(chatID, locale.GetText(lang, "not_subscribed_message"))
		msg.ReplyMarkup = generateKeyboard(lang, channels)
	}

	sendMessage(bot, msg)
}

// The function `processSuccess` sends a success message with an animation to a Telegram chat based on
// the type of update received.
func processSuccess(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var chatID int64

	if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.Message.Chat.ID

		callbackResp := tgbotapi.NewCallback(update.CallbackQuery.ID, "🎉 Доступ открыт!")
		_, err := bot.Request(callbackResp)
		if err != nil {
			logger.Error("❌ Error sending callback: %v", err)
		}
	} else if update.Message != nil {
		chatID = update.Message.Chat.ID
	} else {
		logger.Error("❌ Error: update contains neither Message nor CallbackQuery")
		return
	}

	animation := tgbotapi.NewAnimation(chatID, tgbotapi.FileURL("https://i.imgur.com/MptNzLY.mp4"))

	_, err := bot.Send(animation)
	if err != nil {
		logger.Error("❌ Error sending animation (GIF mode): %v", err)
	}
}
