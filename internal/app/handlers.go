package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleMessage handles messages.
func (a *App) HandleMessage(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		a.HandleStart(update)
	case "check":
		a.HandleCheck(update)
	default:
		a.HandleUnknown(update)
	}
}

// HandleCallbackQuery handles callback queries.
func (a *App) HandleCallbackQuery(update tgbotapi.Update) {
	if update.CallbackQuery == nil {
		return
	}

	switch update.CallbackQuery.Data {
	case "check":
		a.HandleCheck(update)
	case "success":
		a.HandleSuccess(update)
	}
}

// Handle the /start command.
func (a *App) HandleStart(update tgbotapi.Update) {
	lang := a.GetUserLanguage(update)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, a.LocaleStorage.GetText(lang, "start_message"))
	msg.ReplyMarkup = a.GenerateKeyboard(lang)
	a.SendMessage(msg)
}

// Handle the /check command and the "check" callback query.
func (a *App) HandleCheck(update tgbotapi.Update) {
	a.ProcessCheckSubscription(update)
}

// handleUnknown handles unknown commands.
func (a *App) HandleUnknown(update tgbotapi.Update) {
	lang := a.GetUserLanguage(update)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, a.LocaleStorage.GetText(lang, "unknown_command_message"))
	a.SendMessage(msg)
}

// The function `handleSuccess` in Go calls another function `processSuccess` with the provided bot and
// update parameters.
func (a *App) HandleSuccess(update tgbotapi.Update) {
	a.ProcessSuccess(update)
}
