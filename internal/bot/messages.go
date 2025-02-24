package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/pkg/logger"
)

// The `sendMessage` function in Go sets the message parse mode to Markdown and sends the message using
// a Telegram bot API.
func sendMessage(bot *tgbotapi.BotAPI, message tgbotapi.MessageConfig) {
	message.ParseMode = "Markdown"
	_, err := bot.Send(message)
	if err != nil {
		logger.Error("❌ Error sending message: %v", err)
	}
}
