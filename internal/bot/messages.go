package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/internal/app"
)

// The `sendMessage` function in Go sets the message parse mode to Markdown and sends the message using
// a Telegram bot API.
func SendMessage(app *app.App, message tgbotapi.MessageConfig) {
	message.ParseMode = "Markdown"

	_, err := app.Bot.Send(message)
	if err != nil {
		app.Logger.Error("❌ Error sending message: %v", err)
	}
}
