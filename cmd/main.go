package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/config"
	"github.com/vaaxooo/tgaccessbot/internal/app"
	"github.com/vaaxooo/tgaccessbot/internal/bot"
	"github.com/vaaxooo/tgaccessbot/internal/locale"
	"github.com/vaaxooo/tgaccessbot/internal/services"
	"github.com/vaaxooo/tgaccessbot/pkg/logger"
)

func main() {
	// Load configuration
	cfg := config.MustLoadConfig()

	// Initialize logger
	logService := logger.NewLogger(cfg.Debug)

	// Load translations
	localeStorage := locale.NewStorage()
	localeStorage.LoadTranslations()

	// Initialize Telegram bot
	botAPI, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		logService.Error("❌ Error initializing Telegram bot: %v", err)
		panic(err)
	}

	botAPI.Debug = cfg.Debug

	telegramBot := bot.NewTelegramBot(botAPI)

	// Initialize subscription service
	subService := services.NewSubscriptionService(logService, telegramBot, cfg.TelegramChannels)

	// Initialize application
	application := app.NewApp(
		logService,
		telegramBot,
		subService,
		localeStorage,
		cfg.TelegramChannels,
		cfg.SuccessRedirectURL,
	)

	// Start the application
	application.Start()
}
