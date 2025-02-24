package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/internal/app"
)

// Handle the /start command.
func HandleStart(app *app.App, update tgbotapi.Update) {
	lang := GetUserLanguage(update)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, app.LocaleStorage.GetText(lang, "start_message"))
	msg.ReplyMarkup = GenerateKeyboard(app, lang)
	SendMessage(app, msg)
}

// Handle the /check command and the "check" callback query.
func HandleCheck(app *app.App, update tgbotapi.Update) {
	ProcessCheckSubscription(app, update)
}

// handleUnknown handles unknown commands.
func HandleUnknown(app *app.App, update tgbotapi.Update) {
	lang := GetUserLanguage(update)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, app.LocaleStorage.GetText(lang, "unknown_command_message"))
	SendMessage(app, msg)
}

// The function `handleSuccess` in Go calls another function `processSuccess` with the provided bot and
// update parameters.
func HandleSuccess(app *app.App, update tgbotapi.Update) {
	ProcessSuccess(app, update)
}
