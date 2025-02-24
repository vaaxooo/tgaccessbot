package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/internal/locale"
	"github.com/vaaxooo/tgaccessbot/pkg/logger"
)

// The function generates an inline keyboard for a Telegram bot with buttons for subscribing to
// channels and checking something.
func generateKeyboard(lang string, channels []string) tgbotapi.InlineKeyboardMarkup {
	var keyboardRows [][]tgbotapi.InlineKeyboardButton
	for _, channel := range channels {
		keyboardRows = append(keyboardRows, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(locale.GetText(lang, "subscribe_button"), "https://t.me/"+channel),
		))
	}

	keyboardRows = append(keyboardRows, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(locale.GetText(lang, "check_button"), "check"),
	))

	return tgbotapi.NewInlineKeyboardMarkup(keyboardRows...)
}

// The function generates an inline keyboard with a single button for success in the specified
// language.
func generateSuccessKeyboard(lang string, successRedirectURL string) tgbotapi.InlineKeyboardMarkup {
	logger.Debug("Success redirect URL: %s", successRedirectURL)
	keyboardRows := [][]tgbotapi.InlineKeyboardButton{
		// {
		// 	tgbotapi.NewInlineKeyboardButtonData(locale.GetText(lang, "success_button"), "success"),
		// },
		{
			tgbotapi.NewInlineKeyboardButtonURL(locale.GetText(lang, "success_button"), successRedirectURL),
		},
	}
	return tgbotapi.NewInlineKeyboardMarkup(keyboardRows...)
}
