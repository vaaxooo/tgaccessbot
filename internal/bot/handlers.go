package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/internal/locale"
	"github.com/vaaxooo/tgaccessbot/internal/services"
)

// Handle the /start command
func handleStart(bot *tgbotapi.BotAPI, update tgbotapi.Update, channels []string) {
	lang := getUserLanguage(update)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, locale.GetText(lang, "start_message"))
	msg.ReplyMarkup = generateKeyboard(lang, channels)
	sendMessage(bot, msg)
}

// Handle the /check command and the "check" callback query
func handleCheck(bot *tgbotapi.BotAPI, update tgbotapi.Update, subService *services.SubscriptionService, channels []string, successRedirectURL string) {
	processCheckSubscription(bot, update, subService, channels, successRedirectURL)
}

// handleUnknown handles unknown commands
func handleUnknown(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	lang := getUserLanguage(update)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, locale.GetText(lang, "unknown_command_message"))
	sendMessage(bot, msg)
}

// The function `handleSuccess` in Go calls another function `processSuccess` with the provided bot and
// update parameters.
func handleSuccess(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	processSuccess(bot, update)
}
