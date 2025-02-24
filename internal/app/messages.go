package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// The `sendMessage` function in Go sets the message parse mode to Markdown and sends the message using
// a Telegram bot API.
func (a *App) SendMessage(message tgbotapi.MessageConfig) {
	message.ParseMode = "Markdown"

	_, err := a.Bot.Send(message)
	if err != nil {
		a.Logger.Error("❌ Error sending message: %v", err)
	}
}
