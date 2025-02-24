package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// The function generates an inline keyboard for a Telegram bot with buttons for subscribing to
// channels and checking something.
func (a *App) GenerateKeyboard(lang string) tgbotapi.InlineKeyboardMarkup {
	channels := a.TelegramChannels
	keyboardRows := make([][]tgbotapi.InlineKeyboardButton, 0, len(channels)+1)

	for _, channel := range channels {
		keyboardRows = append(keyboardRows, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(a.LocaleStorage.GetText(lang, "subscribe_button"), "https://t.me/"+channel),
		))
	}

	keyboardRows = append(keyboardRows, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(a.LocaleStorage.GetText(lang, "check_button"), "check"),
	))

	return tgbotapi.NewInlineKeyboardMarkup(keyboardRows...)
}

// The function generates an inline keyboard with a single button for success in the specified
// language.
func (a *App) GenerateSuccessKeyboard(lang string) tgbotapi.InlineKeyboardMarkup {
	keyboardRows := [][]tgbotapi.InlineKeyboardButton{
		// {
		// 	tgbotapi.NewInlineKeyboardButtonData(locale.GetText(lang, "success_button"), "success"),
		// },
		{
			tgbotapi.NewInlineKeyboardButtonURL(a.LocaleStorage.GetText(lang, "success_button"), a.SuccessRedirectURL),
		},
	}

	return tgbotapi.NewInlineKeyboardMarkup(keyboardRows...)
}
