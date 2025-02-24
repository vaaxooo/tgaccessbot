package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vaaxooo/tgaccessbot/config"
	"github.com/vaaxooo/tgaccessbot/internal/locale"
	"github.com/vaaxooo/tgaccessbot/internal/services"
	"github.com/vaaxooo/tgaccessbot/pkg/logger"
)

// NewApp creates a new instance of the App structure.
func NewApp() *App {
	// Load the configuration
	config := config.MustLoadConfig()

	// Initialize the logger
	logger := logger.NewLogger(config.Debug)

	// Initialize the locale storage
	localeStorage := locale.NewStorage()
	localeStorage.LoadTranslations()

	// Initialize the Telegram bot
	botAPI, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		logger.Error("❌ Error initializing Telegram bot: %v", err)
		panic(err)
	}

	botAPI.Debug = config.Debug

	// Initialize the subscription service
	subService := services.NewSubscriptionService(logger, botAPI, config.TelegramChannels)

	return &App{
		Logger:             logger,
		Bot:                botAPI,
		Subscription:       subService,
		LocaleStorage:      localeStorage,
		TelegramChannels:   config.TelegramChannels,
		SuccessRedirectURL: config.SuccessRedirectURL,
	}
}
