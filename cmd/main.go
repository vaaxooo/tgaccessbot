package main

import (
	"github.com/vaaxooo/tgaccessbot/config"
	"github.com/vaaxooo/tgaccessbot/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	bot_internal "github.com/vaaxooo/tgaccessbot/internal/bot"
	"github.com/vaaxooo/tgaccessbot/internal/locale"
	"github.com/vaaxooo/tgaccessbot/internal/services"
)

func main() {
	// Init config
	config, err := config.LoadConfig()
	if err != nil {
		panic("❌ An error occurred while loading the .env file")
	}

	// Init logger
	logger.InitLogger(config.Debug)

	// Load locales
	if err := locale.LoadTranslations(); err != nil {
		logger.Error("❌ An error occurred while loading locales: %v", err)
	}

	// Init Telegram bot
	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		panic("❌ An error occurred while initializing the Telegram bot")
	}
	bot.Debug = config.Debug

	// Init subscription service
	subService := services.NewSubscriptionService(bot, config.TelegramChannels)

	// Start the bot
	bot_internal.Start(bot, subService, config.TelegramChannels, config.SuccessRedirectURL)
}
