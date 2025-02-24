package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/internal/app"
)

// The `Start` function in Go initializes a bot, listens for updates, and handles messages and callback
// queries accordingly.
func Start(app *app.App) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := app.Bot.GetUpdatesChan(u)

	app.Logger.Info("🚀 Bot successfully started")

	for update := range updates {
		if update.Message != nil {
			handleMessage(app, update)
		} else if update.CallbackQuery != nil {
			handleCallbackQuery(app, update)
		}
	}
}

// The function `handleMessage` processes different commands in a Telegram bot update by calling
// specific handler functions based on the command.
func handleMessage(app *app.App, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		HandleStart(app, update)
	case "check":
		HandleCheck(app, update)
	default:
		HandleUnknown(app, update)
	}
}

// The `handleCallbackQuery` function in Go processes a callback query by checking its data and calling
// a specific handler function based on the data.
func handleCallbackQuery(app *app.App, update tgbotapi.Update) {
	if update.CallbackQuery == nil {
		return
	}

	switch update.CallbackQuery.Data {
	case "check":
		HandleCheck(app, update)
	case "success":
		HandleSuccess(app, update)
	}
}
