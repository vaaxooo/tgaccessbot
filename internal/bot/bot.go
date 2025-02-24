package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/internal/services"
	"github.com/vaaxooo/tgaccessbot/pkg/logger"
)

// The `Start` function in Go initializes a bot, listens for updates, and handles messages and callback
// queries accordingly.
func Start(bot *tgbotapi.BotAPI, subService *services.SubscriptionService, channels []string, successRedirectURL string) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	logger.Info("🚀 Bot successfully started")

	logger.Debug("Success redirect URL: %s", successRedirectURL)

	for update := range updates {
		if update.Message != nil {
			handleMessage(bot, update, subService, channels, successRedirectURL)
		} else if update.CallbackQuery != nil {
			handleCallbackQuery(bot, update, subService, channels, successRedirectURL)
		}

	}
}

// The function `handleMessage` processes different commands in a Telegram bot update by calling
// specific handler functions based on the command.
func handleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, subService *services.SubscriptionService, channels []string, successRedirectURL string) {
	switch update.Message.Command() {
	case "start":
		handleStart(bot, update, channels)
	case "check":
		handleCheck(bot, update, subService, channels, successRedirectURL)
	default:
		handleUnknown(bot, update)
	}
}

// The `handleCallbackQuery` function in Go processes a callback query by checking its data and calling
// a specific handler function based on the data.
func handleCallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update, subService *services.SubscriptionService, channels []string, successRedirectURL string) {
	if update.CallbackQuery == nil {
		return
	}

	switch update.CallbackQuery.Data {
	case "check":
		handleCheck(bot, update, subService, channels, successRedirectURL)
	case "success":
		handleSuccess(bot, update)
	}
}
